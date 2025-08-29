package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sm sync.Mutex

func LockExercise() {
	fmt.Println("--------锁机制--------")
	/*
		题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
		考察点 ： sync.Mutex 的使用、并发数据安全。
	*/
	//syncMutex()

	/*
		题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
		考察点 ：原子操作、并发数据安全。
	*/
	atomicAdd()
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func syncMutex() {
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				AddCount(&count)
				fmt.Printf("增加 j = %d count = %d\n", j, count)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("统计结束 count= %d", count)
}

func AddCount(p *int) {
	sm.Lock()
	defer sm.Unlock() //将解锁操作压入延迟调用栈，会等待(*p)++执行完结果侯执行，这里不要选择手动释放，可能会出现无法释放的情况
	// 计数器递增
	(*p)++
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func atomicAdd() {
	var wg sync.WaitGroup
	var count int32 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
				fmt.Printf("增加 j = %d count = %d\n", j, count)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("统计结束 count= %d", count)
}
