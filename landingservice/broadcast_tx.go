package landingservice

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ccIisIaIcat/pancakePrediction/common/method"
	"github.com/ccIisIaIcat/pancakePrediction/config"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BroadcastResult 广播结果
type BroadcastResult struct {
	RPC     string        // RPC 端点
	TxHash  string        // 交易哈希
	Success bool          // 是否成功
	Error   error         // 错误信息
	Latency time.Duration // 延迟时间
}

// TxBroadcaster 交易广播器
type TxBroadcaster struct {
	rpcList []string
	timeout time.Duration
}

// NewTxBroadcaster 创建交易广播器
func NewTxBroadcaster(cfg *config.Config, timeout time.Duration) *TxBroadcaster {
	return &TxBroadcaster{
		rpcList: cfg.RPCList,
		timeout: timeout,
	}
}

// BroadcastTx 向多个节点并发发送已签名的交易
func (b *TxBroadcaster) BroadcastTx(signedTx *types.Transaction) []*BroadcastResult {
	results := make([]*BroadcastResult, len(b.rpcList))
	var mu sync.Mutex

	// 使用 ParallelFor2 并发发送到所有节点
	method.ParallelFor2(b.rpcList, func(index int, rpcURL string) error {
		result := b.sendToNode(rpcURL, signedTx)

		mu.Lock()
		results[index] = result
		mu.Unlock()

		return nil // 不中断其他节点的发送
	})

	return results
}

// sendToNode 向单个节点发送交易
func (b *TxBroadcaster) sendToNode(rpcURL string, signedTx *types.Transaction) *BroadcastResult {
	result := &BroadcastResult{
		RPC:    rpcURL,
		TxHash: signedTx.Hash().Hex(),
	}

	startTime := time.Now()

	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), b.timeout)
	defer cancel()

	// 连接到 RPC 节点
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		result.Error = fmt.Errorf("failed to dial: %w", err)
		result.Latency = time.Since(startTime)
		return result
	}
	defer client.Close()

	// 发送交易
	err = client.SendTransaction(ctx, signedTx)
	result.Latency = time.Since(startTime)

	if err != nil {
		result.Error = err
		result.Success = false
	} else {
		result.Success = true
	}

	return result
}

// PrintResults 打印广播结果
func PrintResults(results []*BroadcastResult) {
	successCount := 0
	failCount := 0

	log.Println("\n📡 Broadcast Results:")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	for i, result := range results {
		if result.Success {
			successCount++
			log.Printf("✅ [%d] %s", i, result.RPC)
			log.Printf("   TxHash: %s", result.TxHash)
			log.Printf("   Latency: %v", result.Latency)
		} else {
			failCount++
			log.Printf("❌ [%d] %s", i, result.RPC)
			log.Printf("   Error: %v", result.Error)
			log.Printf("   Latency: %v", result.Latency)
		}
		log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	}

	log.Printf("\n📊 Summary: %d succeeded, %d failed out of %d total\n",
		successCount, failCount, len(results))
}

// GetFastestSuccess 获取最快成功的结果
func GetFastestSuccess(results []*BroadcastResult) *BroadcastResult {
	var fastest *BroadcastResult

	for _, result := range results {
		if result.Success {
			if fastest == nil || result.Latency < fastest.Latency {
				fastest = result
			}
		}
	}

	return fastest
}
