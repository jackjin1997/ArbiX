package service

import (
	"fmt"
	"log"
	"time"
)

// 假设我们用的是假的API数据来测试套利
func ArbitrageOpportunity() {
	// 模拟交易所价格
	priceExchangeA := 100.5
	priceExchangeB := 101.0

	// 计算套利差价
	diff := priceExchangeB - priceExchangeA
	fmt.Printf("Price diff: %.2f\n", diff)

	if diff > 0.5 {
		log.Println("Arbitrage opportunity found! Execute trade!")
	} else {
		log.Println("No arbitrage opportunity.")
	}
}

// 定时执行套利检查
func StartArbitrageMonitor() {
	for {
		ArbitrageOpportunity()
		time.Sleep(10 * time.Second)
	}
}
