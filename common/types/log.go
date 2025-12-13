package types

// EthSubscriptionNotification eth_subscription 订阅通知
type EthSubscriptionNotification struct {
	JSONRPC string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  EthSubscriptionParams `json:"params"`
}

// EthSubscriptionParams 订阅参数
type EthSubscriptionParams struct {
	Subscription string `json:"subscription"`
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
