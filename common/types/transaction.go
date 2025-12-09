package types

// Transaction Bloxroute traceBlocks 流中的交易结构
type Transaction struct {
	From     string            `json:"from"`
	Gas      string            `json:"gas"`
	GasPrice string            `json:"gasPrice"`
	Hash     string            `json:"hash"`
	Input    string            `json:"input"`
	Nonce    string            `json:"nonce"`
	Value    string            `json:"value"`
	V        string            `json:"v"`
	R        string            `json:"r"`
	S        string            `json:"s"`
	To       string            `json:"to"`
	Logs     []Log             `json:"logs"`
	StateDiff map[string]map[string]string `json:"stateDiff"`
}

// Log 交易日志
type Log struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

// TraceBlockNotification Bloxroute traceBlocks 流通知
type TraceBlockNotification struct {
	JSONRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  TraceBlockParams  `json:"params"`
}

// TraceBlockParams traceBlocks 参数
type TraceBlockParams struct {
	Subscription string     `json:"subscription"`
	Result       BlockTrace `json:"result"`
}

// BlockTrace 区块追踪数据
type BlockTrace struct {
	Hash         string        `json:"hash"`
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
	Withdrawals  []interface{} `json:"withdrawals"`
	Uncles       []interface{} `json:"uncles"`
}

// BlockHeader 区块头信息
type BlockHeader struct {
	ParentHash   string `json:"parentHash"`
	Sha3Uncles   string `json:"sha3Uncles"`
	Miner        string `json:"miner"`
	StateRoot    string `json:"stateRoot"`
	TxRoot       string `json:"transactionsRoot"`
	ReceiptsRoot string `json:"receiptsRoot"`
	LogsBloom    string `json:"logsBloom"`
	Difficulty   string `json:"difficulty"`
	Number       string `json:"number"`
	GasLimit     string `json:"gasLimit"`
	GasUsed      string `json:"gasUsed"`
	Timestamp    string `json:"timestamp"`
	ExtraData    string `json:"extraData"`
	MixHash      string `json:"mixHash"`
	Nonce        string `json:"nonce"`
}
