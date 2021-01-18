package main

import (
	"elk-watcher/request"
	"fmt"
	"log"
	"time"
)

func main() {
	var ch chan int
	ticker := time.NewTicker(time.Second * 60)
	go func() {
		for range ticker.C {
			watch()
		}
		ch <- 1
	}()
	<-ch
}

func watch() {
	value := request.GetLogCount()
	if value >= 30 {
		message := fmt.Sprintf("[Warning] QPS 高于阈值. 当前采样为 %d\n", value)
		request.PostDingTalk(message)
		log.Printf(message)
	} else if value > 0 {
		log.Printf("[Info] QPS: %d 正常\n", value)
	}
}
