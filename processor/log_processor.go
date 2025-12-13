package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ccIisIaIcat/pancakePrediction/common/types"
	"github.com/ccIisIaIcat/pancakePrediction/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// LogProcessor 日志处理器,支持多个 WebSocket 源和去重
type LogProcessor struct {
	mu            sync.RWMutex
	txHashCache   map[string]time.Time // 交易 hash -> 过期时间
	cacheExpiry   time.Duration
	contractABI   abi.ABI
	eventCallback func(*types.LogResult, string, map[string]interface{}, string) // 回调函数: (日志, 事件名, 解析后的参数, endpoint)
}

// NewLogProcessor 创建日志处理器
func NewLogProcessor(cacheExpiry time.Duration) (*LogProcessor, error) {
	// 解析合约 ABI
	contractABI, err := abi.JSON(strings.NewReader(contracts.PancakePredictionMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	return &LogProcessor{
		txHashCache:   make(map[string]time.Time),
		cacheExpiry:   cacheExpiry,
		contractABI:   contractABI,
		eventCallback: defaultEventCallback,
	}, nil
}

// SetEventCallback 设置自定义事件回调函数
func (p *LogProcessor) SetEventCallback(callback func(*types.LogResult, string, map[string]interface{}, string)) {
	p.eventCallback = callback
}

// ProcessLogMessage 处理日志消息(带去重)
func (p *LogProcessor) ProcessLogMessage(rawMessage []byte, endpoint string) error {
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
	txHash := logResult.TransactionHash

	// 去重检查
	if p.isDuplicate(txHash) {
		log.Printf("⏭️  [%s] Skipping duplicate transaction: %s", endpoint, txHash)
		return nil
	}

	// 标记为已处理
	p.markProcessed(txHash)

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
func (p *LogProcessor) isDuplicate(txHash string) bool {
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
func (p *LogProcessor) markProcessed(txHash string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.txHashCache[txHash] = time.Now().Add(p.cacheExpiry)
}

// parseEvent 解析事件
func (p *LogProcessor) parseEvent(logResult *types.LogResult) (string, map[string]interface{}, error) {
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
func (p *LogProcessor) StartCleanup(ctx context.Context, interval time.Duration) {
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
func (p *LogProcessor) cleanupExpiredCache() {
	p.mu.Lock()
	defer p.mu.Unlock()

	now := time.Now()
	count := 0
	for hash, expiry := range p.txHashCache {
		if now.After(expiry) {
			delete(p.txHashCache, hash)
			count++
		}
	}

	if count > 0 {
		log.Printf("🧹 Cleaned up %d expired cache entries", count)
	}
}

// defaultEventCallback 默认事件回调函数
func defaultEventCallback(logResult *types.LogResult, eventName string, eventData map[string]interface{}, endpoint string) {
	log.Printf("\n🎯 Event Detected: %s", eventName)
	log.Printf("   📡 Source: %s", endpoint)
	log.Printf("   📍 Contract: %s", logResult.Address)
	log.Printf("   📦 Block: %s", logResult.BlockNumber)
	log.Printf("   🔗 TxHash: %s", logResult.TransactionHash)
	log.Printf("   📊 Event Data:")

	for key, value := range eventData {
		// 格式化不同类型的值
		var formattedValue string
		switch v := value.(type) {
		case *big.Int:
			formattedValue = v.String()
		case common.Address:
			formattedValue = v.Hex()
		case []byte:
			formattedValue = common.Bytes2Hex(v)
		default:
			formattedValue = fmt.Sprintf("%v", v)
		}
		log.Printf("      • %s: %s", key, formattedValue)
	}
	log.Println()
}
