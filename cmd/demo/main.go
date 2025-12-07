package main

import (
	"fmt"

	"github.com/ccIisIaIcat/pancakePrediction/config"
)

func main() {
	// traceBlocks
	config, err := config.LoadConfigFromDefault()
	if err != nil {
		fmt.Println("加载配置文件时出错:", err)
		return
	}
	fmt.Println(config)
}
