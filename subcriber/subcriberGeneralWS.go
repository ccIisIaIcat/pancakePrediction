package subcriber

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type LogSubscriber struct {
	wsURL     string
	conn      *websocket.Conn
	logChan   chan []byte
	address   string
	subID     string
}

// NewLogSubscriber 创建日志订阅器
func NewLogSubscriber(wsURL string, contractAddress string) *LogSubscriber {
	return &LogSubscriber{
		wsURL:   wsURL,
		address: contractAddress,
		logChan: make(chan []byte, 100),
	}
}

// GetLogChan 获取日志通道
func (s *LogSubscriber) GetLogChan() <-chan []byte {
	return s.logChan
}

// Start 启动订阅(带断线重连)
func (s *LogSubscriber) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Log subscriber stopped")
			if s.conn != nil {
				s.conn.Close()
			}
			return
		default:
			if err := s.connectAndSubscribe(); err != nil {
				log.Printf("❌ [%s] Connection failed: %v, retrying in 5s...", s.wsURL, err)
				time.Sleep(5 * time.Second)
				continue
			}
			s.readLoop(ctx)
		}
	}
}

// connectAndSubscribe 连接并订阅
func (s *LogSubscriber) connectAndSubscribe() error {
	// 连接 WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(s.wsURL, nil)
	if err != nil {
		return fmt.Errorf("dial failed: %w", err)
	}
	s.conn = conn
	log.Printf("✅ Connected to %s", s.wsURL)

	// 发送订阅请求
	subscribeReq := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "eth_subscribe",
		"params":  []interface{}{"logs", map[string]interface{}{"address": s.address}},
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
		log.Printf("✅ Subscribed to logs, subscription ID: %s", s.subID)
	}

	return nil
}

// readLoop 读取消息循环
func (s *LogSubscriber) readLoop(ctx context.Context) {
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
				log.Printf("⚠️ [%s] Read error: %v, reconnecting...", s.wsURL, err)
				return
			}

			// 只转发包含 "eth_subscription" 方法的消息(实际日志数据)
			var msg map[string]interface{}
			if err := json.Unmarshal(message, &msg); err == nil {
				if method, ok := msg["method"].(string); ok && method == "eth_subscription" {
					select {
					case s.logChan <- message:
					case <-ctx.Done():
						return
					default:
						log.Println("⚠️ Log channel full, dropping message")
					}
				}
			}
		}
	}
}
