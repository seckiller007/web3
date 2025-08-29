package main

import (
	"fmt"
	"sync"
	"time"
)

func TaskTwoChannel() {
	fmt.Println("-------通道练习-------")
	/*
		题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
		考察点 ：通道的基本使用、协程间通信。
	*/
	//communication()

	/*
		题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
		考察点 ：通道的缓冲机制。
	*/
	bufferChannel()
}

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func communication() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i //进入通道
			fmt.Printf("发送 i = %d\n", i)
			time.Sleep(time.Second * 1)
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("接收到：", v)
		}
	}()
	wg.Wait()
	fmt.Println("main结束")
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func bufferChannel() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
			fmt.Printf("发送 i = %d\n", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println("接收到：", v)
		}
	}()
	wg.Wait()
	fmt.Println("结束")
}
