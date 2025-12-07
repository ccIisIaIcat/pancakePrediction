package subcriber

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/ccIisIaIcat/pancakePrediction/common/types"
	"github.com/ccIisIaIcat/pancakePrediction/config"
	"github.com/gorilla/websocket"
)

type SubcriberBSCBloxroute struct {
	config          config.BloXrouteConfig
	conn            *websocket.Conn
	msgChan         chan []byte
	subscribeMethod string
}

func NewSubcriber(config config.Config) *SubcriberBSCBloxroute {
	s := &SubcriberBSCBloxroute{
		config:  config.BloXroute,
		msgChan: make(chan []byte, 100),
	}
	return s
}

func (s *SubcriberBSCBloxroute) SubcriberService(subcribeMethod string) error {
	if s.conn == nil {
		return fmt.Errorf("connection not established")
	}

	// 保存订阅方法用于重连
	s.subscribeMethod = subcribeMethod

	subscribeParams := map[string]interface{}{
		"blockchain_network": s.config.BlockchainNetwork,
	}
	if len(s.config.Include) > 0 {
		subscribeParams["include"] = s.config.Include
	} else {
		subscribeParams["include"] = []string{}
	}

	request := types.JsonRPCRequest{
		JSONRPC: "2.0",
		ID:      rand.Uint64(),
		Method:  "subscribe",
		Params:  []interface{}{subcribeMethod, subscribeParams},
	}

	err := s.conn.WriteJSON(request)
	if err != nil {
		return fmt.Errorf("failed to send subscribe request: %w", err)
	}

	log.Printf("📤 Sent subscribe request: method=%s, network=%s", subcribeMethod, s.config.BlockchainNetwork)

	// 读取订阅响应
	var response types.JsonRPCResponse
	err = s.conn.ReadJSON(&response)
	if err != nil {
		return fmt.Errorf("failed to read subscribe response: %w", err)
	}

	return nil
}

// keepAlive 保持连接活跃
func (c *SubcriberBSCBloxroute) keepAlive(ctx context.Context) {
	pingTicker := time.NewTicker(15 * time.Second)
	defer pingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-pingTicker.C:
			if err := c.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
				log.Printf("⚠️  Failed to send ping: %v", err)
				return
			}
		}
	}
}

func (s *SubcriberBSCBloxroute) Connect() error {
	if s.config.AuthHeader == "" {
		return fmt.Errorf("bloXroute Auth header is not set")
	}

	// 创建 WebSocket 连接
	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	headers := make(map[string][]string)
	headers["Authorization"] = []string{s.config.AuthHeader}

	log.Printf("Connecting to %s...", s.config.WSEndpoint)
	conn, _, err := dialer.Dial(s.config.WSEndpoint, headers)
	if err != nil {
		return fmt.Errorf("failed to dial ws: %w", err)
	}
	s.conn = conn

	// 设置读取超时,避免永久阻塞
	s.conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	// 设置 pong 处理器,保持连接活跃
	s.conn.SetPongHandler(func(string) error {
		s.conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		return nil
	})

	log.Printf("✅ Connected to bloXroute: %s", s.config.WSEndpoint)

	return nil
}

// GetMsgChan 获取消息通道
func (s *SubcriberBSCBloxroute) GetMsgChan() <-chan []byte {
	return s.msgChan
}

// Start 启动消息接收，包含断线重连机制
func (s *SubcriberBSCBloxroute) Start(ctx context.Context) {
	go s.keepAlive(ctx)

	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Subscriber context cancelled, stopping...")
			return
		default:
			// 读取消息
			_, message, err := s.conn.ReadMessage()
			if err != nil {
				log.Printf("⚠️  Connection error: %v, reconnecting...", err)

				// 断线重连
				if err := s.reconnectAndResubscribe(ctx); err != nil {
					log.Printf("❌ Failed to reconnect: %v, retrying in 5s...", err)
					time.Sleep(1 * time.Second)
					continue
				}
				continue
			}

			// 发送消息到通道
			select {
			case s.msgChan <- message:
				// 消息成功发送
			case <-ctx.Done():
				return
			default:
				log.Println("⚠️  Message channel full, dropping message")
			}
		}
	}
}

// reconnectAndResubscribe 断线重连并重新订阅
func (s *SubcriberBSCBloxroute) reconnectAndResubscribe(ctx context.Context) error {
	// 关闭旧连接
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}

	// 重新连接
	if err := s.Connect(); err != nil {
		return fmt.Errorf("failed to reconnect: %w", err)
	}

	// 如果有订阅方法，重新订阅
	if s.subscribeMethod != "" {
		if err := s.SubcriberService(s.subscribeMethod); err != nil {
			return fmt.Errorf("failed to resubscribe: %w", err)
		}
		log.Printf("✅ Reconnected and resubscribed to %s", s.subscribeMethod)
	}

	return nil
}
