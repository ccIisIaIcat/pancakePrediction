package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BloXroute BloXrouteConfig `json:"bloxroute"`
}

type BloXrouteConfig struct {
	WSEndpoint        string   `json:"WSEndpoint"`        // bloXroute WebSocket 端点 (例如: wss://mev.api.blxrbdn.com/ws)
	AuthHeader        string   `json:"AuthHeader"`        // 认证令牌
	BlockchainNetwork string   `json:"BlockchainNetwork"` // 区块链网络 (例如: BSC-Mainnet)
	Include           []string `json:"Include"`           // 要包含的字段(可选)
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
