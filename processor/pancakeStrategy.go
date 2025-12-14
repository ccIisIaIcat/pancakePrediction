package processor

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ccIisIaIcat/pancakePrediction/common/types"
	"github.com/ccIisIaIcat/pancakePrediction/config"
	"github.com/ccIisIaIcat/pancakePrediction/contracts"
	"github.com/ccIisIaIcat/pancakePrediction/mail"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogProcessor 日志处理器,支持多个 WebSocket 源和去重
type PancakeStrategy struct {
	mu            sync.RWMutex
	txHashCache   map[string]time.Time // 交易 hash -> 过期时间
	cacheExpiry   time.Duration
	contractABI   abi.ABI
	eventCallback func(*types.LogResult, string, map[string]interface{}, string) // 回调函数: (日志, 事件名, 解析后的参数, endpoint)
	logger        *zap.Logger                                                    // zap 日志对象
	currentDate   string                                                         // 当前日期
	currentBlock  uint64                                                         // 当前区块号

	// 策略状态
	rounds       map[uint64]*RoundState // epoch -> RoundState
	currentEpoch uint64                 // 当前活跃的epoch

	// 策略配置
	config      *config.StrategyConfig
	privateKey  *ecdsa.PrivateKey // 私钥(用于签名交易)
	myAddress   common.Address    // 我的地址
	nonce       uint64            // 当前 nonce（缓存）
	rpcURL      string            // RPC URL（用于获取 nonce）
	rpcList     []string          // 所有 RPC URL列表（用于发送交易）
	riskManager *RiskManager      // 风控管理器
	mailSender  *mail.MailSender  // 邮件发送器
	mailTo      []string          // 邮件接收者列表
}

// NewLogProcessor 创建日志处理器
func NewPancakeStrategy(cacheExpiry time.Duration, strategyConfig *config.StrategyConfig, privateKeyHex string, rpcURL string, rpcList []string, mailConfig *config.MailConfig) (*PancakeStrategy, error) {
	// 解析合约 ABI
	contractABI, err := abi.JSON(strings.NewReader(contracts.PancakePredictionMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// 获取地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to get public key")
	}
	myAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 创建风控管理器
	riskManager, err := NewRiskManager(
		strategyConfig.RiskControl.MaxBetAmount,
		strategyConfig.RiskControl.MinBetAmount,
		strategyConfig.RiskControl.MaxTotalBets,
		strategyConfig.RiskControl.MaxConcurrentBets,
		strategyConfig.RiskControl.MinBalance,
		strategyConfig.RiskControl.StopLoss,
		strategyConfig.RiskControl.DailyLossLimit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create risk manager: %w", err)
	}

	// 创建邮件发送器
	var mailSender *mail.MailSender
	var mailTo []string
	if mailConfig != nil && mailConfig.From != "" && mailConfig.AuthCode != "" {
		mailSender = mail.NewMailSender(
			"smtp.qq.com",
			465,
			mailConfig.From,
			mailConfig.AuthCode,
			"Pancake Strategy Bot",
		)
		mailTo = mailConfig.To
		log.Printf("📧 Mail notification enabled: %s -> %v", mailConfig.From, mailTo)
	} else {
		log.Printf("⚠️ Mail notification disabled (no mail config)")
	}

	p := &PancakeStrategy{
		txHashCache: make(map[string]time.Time),
		cacheExpiry: cacheExpiry,
		contractABI: contractABI,
		rounds:      make(map[uint64]*RoundState),
		config:      strategyConfig,
		privateKey:  privateKey,
		myAddress:   myAddress,
		rpcURL:      rpcURL,
		rpcList:     rpcList,
		nonce:       0, // 初始化时会更新
		riskManager: riskManager,
		mailSender:  mailSender,
		mailTo:      mailTo,
	}
	p.eventCallback = p.defaultEventCallback

	// 初始化 zap 日志
	if err := p.initLogger(); err != nil {
		return nil, err
	}

	log.Printf("📋 Strategy Config: minRatio=%.2f, k=%.2f, blocksPerRound=%d, triggerBlockDiff=%d",
		strategyConfig.MinRatio, strategyConfig.KFactor, strategyConfig.BlocksPerRound, strategyConfig.TriggerBlockDiff)
	log.Printf("🔑 Wallet Address: %s", myAddress.Hex())

	// 初始化 nonce
	if err := p.refreshNonce(); err != nil {
		log.Printf("⚠️ Failed to initialize nonce: %v (will retry)", err)
	}

	return p, nil
}

// GetConfig 获取策略配置（线程安全）
func (p *PancakeStrategy) GetConfig() *config.StrategyConfig {
	return p.config
}

// refreshNonce 从 RPC 获取最新 nonce（只在比当前 nonce 大时更新）
func (p *PancakeStrategy) refreshNonce() error {
	// 导入需要的包（在文件顶部已导入）
	client, err := ethclient.Dial(p.rpcURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer client.Close()

	newNonce, err := client.PendingNonceAt(context.Background(), p.myAddress)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	// 只在新 nonce 更大时更新（保险起见）
	if newNonce > p.nonce {
		oldNonce := p.nonce
		p.nonce = newNonce
		log.Printf("🔄 Nonce updated: %d -> %d", oldNonce, newNonce)
	}

	return nil
}

// incrementNonce 递增 nonce（发送交易后调用）
func (p *PancakeStrategy) incrementNonce() {
	p.nonce++
	log.Printf("➕ Nonce incremented: %d", p.nonce)
}

// getCurrentNonce 获取当前 nonce（用于构造交易）
func (p *PancakeStrategy) getCurrentNonce() uint64 {
	return p.nonce
}

// initLogger 初始化 zap 日志
func (p *PancakeStrategy) initLogger() error {
	today := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("logprocess_%s.log", today)
	p.currentDate = today

	// 配置 zap
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	config.EncodeLevel = zapcore.LowercaseLevelEncoder

	fileEncoder := zapcore.NewJSONEncoder(config)

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	writer := zapcore.AddSync(logFile)
	core := zapcore.NewCore(fileEncoder, writer, zapcore.InfoLevel)
	p.logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Named("log_processor")

	return nil
}

// checkAndRotateLogger 检查并轮换日志
func (p *PancakeStrategy) checkAndRotateLogger() {
	today := time.Now().Format("2006-01-02")
	if p.currentDate != today {
		p.mu.Lock()
		defer p.mu.Unlock()

		if p.logger != nil {
			p.logger.Sync()
		}
		p.initLogger()
	}
}

// SetEventCallback 设置自定义事件回调函数
func (p *PancakeStrategy) SetEventCallback(callback func(*types.LogResult, string, map[string]interface{}, string)) {
	p.eventCallback = callback
}

// ProcessLogMessage 处理日志消息(带去重)
func (p *PancakeStrategy) ProcessLogMessage(rawMessage []byte, endpoint string) error {
	// 解析消息
	var notification types.EthSubscriptionNotification
	if err := json.Unmarshal(rawMessage, &notification); err != nil {
		return fmt.Errorf("failed to unmarshal notification: %w", err)
	}

	// 检查是否为 eth_subscription 消息
	if notification.Method != "eth_subscription" {
		return nil
	}

	logResult := &notification.Params.Result
	// 使用 txHash + logIndex 作为唯一标识
	logKey := logResult.TransactionHash + ":" + logResult.LogIndex

	// 去重检查
	if p.isDuplicate(logKey) {
		return nil
	}

	// 标记为已处理
	p.markProcessed(logKey)

	// 解析事件
	eventName, eventData, err := p.parseEvent(logResult)
	if err != nil {
		log.Printf("⚠️  [%s] Failed to parse event: %v", endpoint, err)
		return err
	}

	// 调用回调函数
	if p.eventCallback != nil {
		p.eventCallback(logResult, eventName, eventData, endpoint)
	}

	return nil
}

// isDuplicate 检查交易是否已处理(带过期清理)
func (p *PancakeStrategy) isDuplicate(txHash string) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 清理过期缓存
	now := time.Now()
	for hash, expiry := range p.txHashCache {
		if now.After(expiry) {
			delete(p.txHashCache, hash)
		}
	}

	// 检查是否存在
	_, exists := p.txHashCache[txHash]
	return exists
}

// markProcessed 标记交易为已处理
func (p *PancakeStrategy) markProcessed(txHash string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.txHashCache[txHash] = time.Now().Add(p.cacheExpiry)
}

// parseEvent 解析事件
func (p *PancakeStrategy) parseEvent(logResult *types.LogResult) (string, map[string]interface{}, error) {
	if len(logResult.Topics) == 0 {
		return "", nil, fmt.Errorf("no topics in log")
	}

	// 构造 go-ethereum 的 Log 类型
	topics := make([]common.Hash, len(logResult.Topics))
	for i, topic := range logResult.Topics {
		topics[i] = common.HexToHash(topic)
	}

	ethLog := ethtypes.Log{
		Address: common.HexToAddress(logResult.Address),
		Topics:  topics,
		Data:    common.FromHex(logResult.Data),
	}

	// 根据 topic[0] 查找事件
	eventSig := topics[0]
	event, err := p.contractABI.EventByID(eventSig)
	if err != nil {
		return "", nil, fmt.Errorf("unknown event signature %s: %w", eventSig.Hex(), err)
	}

	// 解析事件数据
	eventData := make(map[string]interface{})
	if err := p.contractABI.UnpackIntoMap(eventData, event.Name, ethLog.Data); err != nil {
		return "", nil, fmt.Errorf("failed to unpack event data: %w", err)
	}

	// 解析索引参数(topics)
	if len(topics) > 1 {
		indexedArgs := make([]abi.Argument, 0)
		for _, arg := range event.Inputs {
			if arg.Indexed {
				indexedArgs = append(indexedArgs, arg)
			}
		}

		if err := abi.ParseTopicsIntoMap(eventData, indexedArgs, topics[1:]); err != nil {
			return "", nil, fmt.Errorf("failed to parse indexed arguments: %w", err)
		}
	}

	return event.Name, eventData, nil
}

// StartCleanup 启动定期清理过期缓存
func (p *PancakeStrategy) StartCleanup(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p.cleanupExpiredCache()
		}
	}
}

// cleanupExpiredCache 清理过期缓存
func (p *PancakeStrategy) cleanupExpiredCache() {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 清理 txHash 缓存
	now := time.Now()
	count := 0
	for hash, expiry := range p.txHashCache {
		if now.After(expiry) {
			delete(p.txHashCache, hash)
			count++
		}
	}

	if count > 0 {
		log.Printf("🧹 Cleaned up %d expired tx cache entries", count)
	}

	// 清理旧轮次，只保留最新的20个
	if len(p.rounds) > 20 {
		// 收集所有 epoch 并排序
		epochs := make([]uint64, 0, len(p.rounds))
		for epoch := range p.rounds {
			epochs = append(epochs, epoch)
		}

		// 排序（从小到大）
		for i := 0; i < len(epochs)-1; i++ {
			for j := i + 1; j < len(epochs); j++ {
				if epochs[i] > epochs[j] {
					epochs[i], epochs[j] = epochs[j], epochs[i]
				}
			}
		}

		// 删除最旧的轮次（保留最新20个）
		toDelete := len(epochs) - 20
		deletedCount := 0
		for i := 0; i < toDelete; i++ {
			delete(p.rounds, epochs[i])
			deletedCount++
		}

		if deletedCount > 0 {
			log.Printf("🧹 Cleaned up %d old rounds, keeping latest 20", deletedCount)
		}
	}
}

// UpdateBlockNumber 更新当前区块号(仅当新区块号更大时)
func (p *PancakeStrategy) UpdateBlockNumber(blockNumber uint64) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	if blockNumber > p.currentBlock {
		p.currentBlock = blockNumber

		// 检查所有活跃轮次的下注时机
		p.checkBetOpportunities(blockNumber)

		return true
	}
	return false
}

// GetCurrentBlock 获取当前区块号
func (p *PancakeStrategy) GetCurrentBlock() uint64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.currentBlock
}

// GetRound 获取指定 epoch 的 RoundState（线程安全）
func (p *PancakeStrategy) GetRound(epoch uint64) (*RoundState, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	round, exists := p.rounds[epoch]
	return round, exists
}

// GetCurrentEpoch 获取当前活跃的 epoch（线程安全）
func (p *PancakeStrategy) GetCurrentEpoch() uint64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.currentEpoch
}

// defaultEventCallback 默认事件回调函数
func (p *PancakeStrategy) defaultEventCallback(logResult *types.LogResult, eventName string, eventData map[string]interface{}, endpoint string) {
	p.checkAndRotateLogger()

	// 事件分发到对应的处理器
	switch eventName {
	case "StartRound":
		p.handleStartRound(logResult, eventData)
	case "BetBull":
		p.handleBetBull(logResult, eventData)
	case "BetBear":
		p.handleBetBear(logResult, eventData)
	case "LockRound":
		p.handleLockRound(logResult, eventData)
	case "EndRound":
		p.handleEndRound(logResult, eventData)
	case "Claim":
		p.handleClaim(logResult, eventData)
	}

	// 记录日志
	p.logEvent(logResult, eventName, eventData, endpoint)
}

// logEvent 记录事件日志
func (p *PancakeStrategy) logEvent(logResult *types.LogResult, eventName string, eventData map[string]interface{}, endpoint string) {
	// 构建 zap 字段
	fields := []zap.Field{
		zap.String("endpoint", endpoint),
		zap.String("contract", logResult.Address),
		zap.String("blockNumber", logResult.BlockNumber),
		zap.String("txHash", logResult.TransactionHash),
		zap.String("blockHash", logResult.BlockHash),
		zap.String("logIndex", logResult.LogIndex),
	}

	// 添加事件数据字段
	for key, value := range eventData {
		switch v := value.(type) {
		case *big.Int:
			fields = append(fields, zap.String(key, v.String()))
		case common.Address:
			fields = append(fields, zap.String(key, v.Hex()))
		case []byte:
			fields = append(fields, zap.String(key, common.Bytes2Hex(v)))
		default:
			fields = append(fields, zap.Any(key, v))
		}
	}

	p.logger.Info("Event Detected",
		append([]zap.Field{zap.String("eventName", eventName)}, fields...)...)
}

