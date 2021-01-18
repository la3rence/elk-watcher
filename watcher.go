package main

import (
	"elk-watcher/request"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var ch chan int
	durationString := os.Getenv("DURATION")
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		log.Fatal("Wrong params: DURATION")
	}
	ticker := time.NewTicker(duration)
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
