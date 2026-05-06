package main

import (
	"staking-interaction/listener"
)

func main() {
	// 启动事件监听器
	go listener.StartStakeEventListener()
}
