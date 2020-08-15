package main

import (
	"fmt"
	"time"
)

var i int = 0

func timeMore(ch chan string) {

	// 执行前先注册，写不进去就会阻塞
	ch <- "任务"
	i++
	fmt.Println("模拟耗时操作", i)
	time.Sleep(time.Second) // 模拟耗时操作

	// 任务执行完毕，则管道中销毁一个任务
	<-ch
}

func main() {

	ch := make(chan string, 5)

	// 开启100个协程
	for i := 0; i < 10; i++ {
		go timeMore(ch)
	}

	for {
		time.Sleep(time.Second * 5)
	}
}