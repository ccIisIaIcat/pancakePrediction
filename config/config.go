package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ChainID       int64           `json:"chainID"`
	BloXroute     BloXrouteConfig `json:"bloxroute"`
	WebsocketList []string        `json:"wslist"`
	RPCList       []string        `json:"rpcList"`
	Strategy      StrategyConfig  `json:"strategy"`
	Mail          MailConfig      `json:"mail"`
}

type BloXrouteConfig struct {
	WSEndpoint        string   `json:"WSEndpoint"`        // bloXroute WebSocket 端点 (例如: wss://mev.api.blxrbdn.com/ws)
	AuthHeader        string   `json:"AuthHeader"`        // 认证令牌
	BlockchainNetwork string   `json:"BlockchainNetwork"` // 区块链网络 (例如: BSC-Mainnet)
	Include           []string `json:"Include"`           // 要包含的字段(可选)
}

// StrategyConfig 策略配置
type StrategyConfig struct {
	ContractAddress  string       `json:"contractAddress"`  // 合约地址
	MinRatio         float64      `json:"minRatio"`         // 最小ratio阈值 (例如: 2.0)
	KFactor          float64      `json:"kFactor"`          // 下注系数 (例如: 0.10, 0.15)
	BlocksPerRound   uint64       `json:"blocksPerRound"`   // 每轮区块数 (例如: 508)
	TriggerBlockDiff uint64       `json:"triggerBlockDiff"` // 提前几个区块下注 (例如: 7-8)
	GasPrice         int64        `json:"gasPrice"`         // Gas价格 (Gwei, 例如: 3000000000 = 3 Gwei)
	GasLimitBet      uint64       `json:"gasLimitBet"`      // 下注Gas限制 (例如: 200000)
	GasLimitClaim    uint64       `json:"gasLimitClaim"`    // Claim Gas限制 (例如: 250000)
	RiskControl      RiskControl  `json:"riskControl"`      // 风控配置
}

// RiskControl 风控配置
type RiskControl struct {
	MaxBetAmount       string `json:"maxBetAmount"`       // 单笔最大下注金额 (wei, 例如: "1000000000000000000" = 1 BNB)
	MinBetAmount       string `json:"minBetAmount"`       // 单笔最小下注金额 (wei, 例如: "10000000000000000" = 0.01 BNB)
	MaxTotalBets       int    `json:"maxTotalBets"`       // 最大下注次数 (0 = 无限制)
	MaxConcurrentBets  int    `json:"maxConcurrentBets"`  // 最大并发下注数 (同时未结束的下注, 0 = 无限制)
	MinBalance         string `json:"minBalance"`         // 最小余额保护 (wei)
	StopLoss           string `json:"stopLoss"`           // 止损金额 (wei, 0 = 不启用)
	DailyLossLimit     string `json:"dailyLossLimit"`     // 每日亏损限制 (wei, 0 = 不启用)
}

// MailConfig 邮箱配置
type MailConfig struct {
	AuthCode string   `json:"AuthCode"` // 邮箱授权码
	From     string   `json:"From"`     // 发件人邮箱
	To       []string `json:"To"`       // 收件人邮箱列表
}

// LoadConfig 从指定的 JSON 文件加载配置
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法打开配置文件: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("无法解析配置文件: %w", err)
	}

	return &config, nil
}

// LoadConfigFromDefault 从默认路径加载配置文件
func LoadConfigFromDefault() (*Config, error) {
	return LoadConfig("config.json")
}
