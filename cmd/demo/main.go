package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ccIisIaIcat/pancakePrediction/config"
	"github.com/ccIisIaIcat/pancakePrediction/subcriber"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfigFromDefault()
	if err != nil {
		log.Fatalf("❌ 加载配置文件时出错: %v", err)
	}
	log.Printf("✅ 配置加载成功: %+v\n", cfg.BloXroute)

	// 创建订阅器
	subscriber := subcriber.NewSubcriber(*cfg)
	log.Println("✅ 订阅器创建成功")

	// 连接到 Bloxroute
	err = subscriber.Connect()
	if err != nil {
		log.Fatalf("❌ 连接 Bloxroute 失败: %v", err)
	}
	log.Println("✅ 已连接到 Bloxroute")

	// 订阅 newTxs 流（新交易）
	err = subscriber.SubcriberService("traceBlocks")
	if err != nil {
		log.Fatalf("❌ 订阅 newTxs 失败: %v", err)
	}
	log.Println("✅ 已订阅 newTxs 流")

	// 创建带取消功能的 context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动消息接收（在后台 goroutine 中）
	go subscriber.Start(ctx)
	log.Println("🚀 开始接收消息...")

	// 监听消息通道并打印
	msgChan := subscriber.GetMsgChan()

	// 设置信号处理，允许优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log.Println("📡 等待交易流数据... (按 Ctrl+C 退出)")

	for {
		select {
		case msg := <-msgChan:
			// 打印接收到的消息
			fmt.Printf("\n📨 收到交易流数据:\n%s\n", string(msg))
			fmt.Println("----------------------------------------")
		case <-sigChan:
			log.Println("\n🛑 收到退出信号，正在关闭...")
			cancel()
			return
		}
	}
}
