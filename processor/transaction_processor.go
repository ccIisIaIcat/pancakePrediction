package processor

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/ccIisIaIcat/pancakePrediction/common/types"
	"github.com/ccIisIaIcat/pancakePrediction/subcriber"
)

// TransactionCallback 交易回调函数类型
type TransactionCallback func(tx *types.Transaction, blockNumber string, blockHash string, timestamp string)

// TransactionProcessor 交易处理器
type TransactionProcessor struct {
	subscriber *subcriber.SubcriberBSCBloxroute
	callbacks  []TransactionCallback
	mu         sync.RWMutex
}

// NewTransactionProcessor 创建新的交易处理器
func NewTransactionProcessor(subscriber *subcriber.SubcriberBSCBloxroute) *TransactionProcessor {
	return &TransactionProcessor{
		subscriber: subscriber,
		callbacks:  make([]TransactionCallback, 0),
	}
}

// RegisterCallback 注册交易回调函数
func (p *TransactionProcessor) RegisterCallback(callback TransactionCallback) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.callbacks = append(p.callbacks, callback)
	log.Printf("✅ 注册交易回调函数，当前共 %d 个回调", len(p.callbacks))
}

// ClearCallbacks 清空所有回调函数
func (p *TransactionProcessor) ClearCallbacks() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.callbacks = make([]TransactionCallback, 0)
	log.Println("🗑️  已清空所有交易回调函数")
}

// executeCallbacks 执行所有注册的回调函数
func (p *TransactionProcessor) executeCallbacks(tx *types.Transaction, blockNumber, blockHash, timestamp string) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for i, callback := range p.callbacks {
		// 在独立的 goroutine 中执行回调，避免阻塞
		go func(idx int, cb TransactionCallback) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("⚠️  回调函数 #%d 发生 panic: %v", idx, r)
				}
			}()
			cb(tx, blockNumber, blockHash, timestamp)
		}(i, callback)
	}
}

// Start 启动交易处理器
func (p *TransactionProcessor) Start(ctx context.Context) {
	log.Println("🚀 交易处理器启动中...")

	// 获取消息通道
	msgChan := p.subscriber.GetMsgChan()

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 交易处理器已停止")
			return

		case rawMsg := <-msgChan:
			// 解析消息
			p.processMessage(rawMsg)
		}
	}
}

// processMessage 处理接收到的消息
func (p *TransactionProcessor) processMessage(rawMsg []byte) {
	// 解析为 TraceBlockNotification
	var notification types.TraceBlockNotification
	if err := json.Unmarshal(rawMsg, &notification); err != nil {
		log.Printf("⚠️  解析消息失败: %v", err)
		return
	}

	// 检查是否是 subscribe 方法的通知
	if notification.Method != "subscribe" {
		return
	}

	blockTrace := notification.Params.Result
	blockNumber := blockTrace.Header.Number
	blockHash := blockTrace.Hash
	timestamp := blockTrace.Header.Timestamp

	// 处理区块中的每一笔交易
	if len(blockTrace.Transactions) == 0 {
		return
	}

	log.Printf("📦 区块 %s 包含 %d 笔交易", blockNumber, len(blockTrace.Transactions))

	for i := range blockTrace.Transactions {
		tx := &blockTrace.Transactions[i]

		// 执行所有注册的回调函数
		p.executeCallbacks(tx, blockNumber, blockHash, timestamp)
	}
}
