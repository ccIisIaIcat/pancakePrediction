package subcriber

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

// BlockNumberData 区块号数据
type BlockNumberData struct {
	Number uint64 // 区块号(已转换为十进制)
	Hex    string // 十六进制原始值
}

// BlockSubscriber 区块订阅器
type BlockSubscriber struct {
	wsURL      string
	conn       *websocket.Conn
	blockChan  chan *BlockNumberData
	subID      string
}

// NewBlockSubscriber 创建区块订阅器
func NewBlockSubscriber(wsURL string) *BlockSubscriber {
	return &BlockSubscriber{
		wsURL:     wsURL,
		blockChan: make(chan *BlockNumberData, 100),
	}
}

// GetBlockChan 获取区块通道
func (s *BlockSubscriber) GetBlockChan() <-chan *BlockNumberData {
	return s.blockChan
}

// Start 启动订阅(带断线重连)
func (s *BlockSubscriber) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Block subscriber stopped")
			if s.conn != nil {
				s.conn.Close()
			}
			return
		default:
			if err := s.connectAndSubscribe(); err != nil {
				log.Printf("❌ [%s] Block connection failed: %v, retrying in 5s...", s.wsURL, err)
				time.Sleep(5 * time.Second)
				continue
			}
			s.readLoop(ctx)
		}
	}
}

// connectAndSubscribe 连接并订阅
func (s *BlockSubscriber) connectAndSubscribe() error {
	// 连接 WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(s.wsURL, nil)
	if err != nil {
		return fmt.Errorf("dial failed: %w", err)
	}
	s.conn = conn
	log.Printf("✅ Block subscriber connected to %s", s.wsURL)

	// 发送订阅请求 - 订阅新区块头
	subscribeReq := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_subscribe",
		"params":  []interface{}{"newHeads"},
	}

	if err := s.conn.WriteJSON(subscribeReq); err != nil {
		s.conn.Close()
		return fmt.Errorf("subscribe failed: %w", err)
	}

	// 读取订阅响应获取 subscription ID
	var response map[string]interface{}
	if err := s.conn.ReadJSON(&response); err != nil {
		s.conn.Close()
		return fmt.Errorf("read subscribe response failed: %w", err)
	}

	if result, ok := response["result"].(string); ok {
		s.subID = result
		log.Printf("✅ Subscribed to newHeads, subscription ID: %s", s.subID)
	}

	return nil
}

// readLoop 读取消息循环
func (s *BlockSubscriber) readLoop(ctx context.Context) {
	defer func() {
		if s.conn != nil {
			s.conn.Close()
			s.conn = nil
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := s.conn.ReadMessage()
			if err != nil {
				log.Printf("⚠️ [%s] Block read error: %v, reconnecting...", s.wsURL, err)
				return
			}

			// 解析消息
			var msg map[string]interface{}
			if err := json.Unmarshal(message, &msg); err != nil {
				continue
			}

			// 只处理 eth_subscription 消息
			if method, ok := msg["method"].(string); ok && method == "eth_subscription" {
				if params, ok := msg["params"].(map[string]interface{}); ok {
					if result, ok := params["result"].(map[string]interface{}); ok {
						if numberHex, ok := result["number"].(string); ok {
							// 转换十六进制区块号为十进制
							blockNum, err := strconv.ParseUint(numberHex[2:], 16, 64) // 去掉 "0x" 前缀
							if err != nil {
								log.Printf("⚠️ Failed to parse block number %s: %v", numberHex, err)
								continue
							}

							blockData := &BlockNumberData{
								Number: blockNum,
								Hex:    numberHex,
							}

							select {
							case s.blockChan <- blockData:
							case <-ctx.Done():
								return
							default:
								log.Println("⚠️ Block channel full, dropping message")
							}
						}
					}
				}
			}
		}
	}
}
