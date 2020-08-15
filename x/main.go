package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)
var wg sync.WaitGroup

func timeMore(ch chan string,i int)  {
	
	// 执行前先注册，写不进去就会阻塞
	ch <- "任务"+strconv.Itoa(i)

	fmt.Println("模拟耗时操作")
	time.Sleep(time.Second)	// 模拟耗时操作

	// 任务执行完毕，则管道中销毁一个任务
	str1<-ch
	wg.Done()
	// return str1
}

func main() {


	ch := make(chan string, 5)
	// s1 := []string{}

	// 开启100个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		timeMore(ch,i)
		s1= append(s1,s)
	}

	wg.Wait()
	fmt.Println(s1)

}