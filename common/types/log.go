package types

// EthSubscriptionNotification eth_subscription 订阅通知
type EthSubscriptionNotification struct {
	JSONRPC string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  EthSubscriptionParams `json:"params"`
}

// EthSubscriptionParams 订阅参数
type EthSubscriptionParams struct {
	Subscription string    `json:"subscription"`
	Result       LogResult `json:"result"`
}

// LogResult 日志结果
type LogResult struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	BlockTimestamp   string   `json:"blockTimestamp,omitempty"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

// BlockHeaderNotification 区块头订阅通知
type BlockHeaderNotification struct {
	JSONRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  BlockHeaderParams `json:"params"`
}

// BlockHeaderParams 区块头订阅参数
type BlockHeaderParams struct {
	Subscription string      `json:"subscription"`
	Result       BlockHeader `json:"result"`
}

// // BlockHeader 区块头(仅包含必要字段)
// type BlockHeader struct {
// 	Number string `json:"number"` // 十六进制区块号
// 	Hash   string `json:"hash"`   // 区块哈希
// }
