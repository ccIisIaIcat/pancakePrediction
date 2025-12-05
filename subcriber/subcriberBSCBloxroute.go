package subcriber

import (
	"fmt"
	"log"
	"time"

	"github.com/ccIisIaIcat/pancakePrediction/config"
	"github.com/gorilla/websocket"
)

type SubcriberBSCBloxroute struct {
	config config.BloXrouteConfig
	conn   *websocket.Conn
}

func NewSubcriber(config config.Config) *SubcriberBSCBloxroute {
	s := &SubcriberBSCBloxroute{
		config: config.BloXroute,
	}
	return s
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
