package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutine
func Goroutine() {
	fmt.Println("-----goroutine-----")
	/*
		题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
		考察点 ： go 关键字的使用、协程的并发执行。
	*/
	SubThread()

	/*
		题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
		考察点 ：协程原理、并发任务调度。
	*/
	//TaskScheduler()

}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func SubThread() {
	fmt.Println("-----sub thread-----")
	var wg sync.WaitGroup //协调多个goroutine同步执行
	wg.Add(2)
	// 打印奇数
	go func() {
		defer wg.Done() //执行完这个goroutine计数器减1
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Printf("打印奇数 i = %d\n", i)
			}
		}
	}()
	// 打印偶数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Printf("打印偶数 i = %d\n", i)
			}
		}
	}()
	wg.Wait() //主goroutine调用Wait()时会阻塞，直到计数器归零，没有控制顺序，打印出是无序的
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 定义调度器
type Scheduler struct {
	Name     string         // 调度器名字
	taskList []*Task        // 任务列表
	wg       sync.WaitGroup // 等待组，用于等待全部任务执行完成
}

// 定义任务
type Task struct {
	ID       int           // 任务id
	Name     string        // 任务名字
	Func     func()        // 任务函数
	Duration time.Duration // 任务执行时间
}

// 创建调度器
func NewScheduler(name string) *Scheduler {
	taskList := []*Task{} //初始化一个空的切片
	return &Scheduler{name, taskList, sync.WaitGroup{}}
}

// 添加任务
func (s *Scheduler) AddTask(task *Task) {
	// 追加任务时, 自定义任务id
	if task.ID == 0 {
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		task.ID = rand.Int()
	}
	s.taskList = append(s.taskList, task)
}

// 运行调度器, 返回结果
func (s *Scheduler) Run() []*Task {
	for _, task := range s.taskList {
		s.wg.Add(1)
		go func(task *Task) {
			defer s.wg.Done()
			start := time.Now()
			task.Func()
			(*task).Duration = time.Since(start)
		}(task)
	}
	s.wg.Wait()
	return s.taskList
}

// 测试用例
func TaskScheduler() {
	fmt.Println("-----task scheduler-----")
	// 创建任务调度器
	scheduler := NewScheduler("任务调度器")
	// 添加任务
	scheduler.AddTask(&Task{
		ID:   123,
		Name: "任务1",
		Func: func() {
			fmt.Println("任务1 执行开始")
			time.Sleep(time.Second * 5)
			fmt.Println("任务1 执行结束")
		},
	})
	scheduler.AddTask(&Task{
		Name: "任务2",
		Func: func() {
			fmt.Println("任务2 执行开始")
			time.Sleep(time.Second * 3)
			fmt.Println("任务2 执行结束")
		},
	})
	scheduler.AddTask(&Task{
		Name: "任务3",
		Func: func() {
			fmt.Println("任务3 执行开始")
			time.Sleep(time.Second * 2)
			fmt.Println("任务3 执行结束")
		},
	})
	// 运行调度器
	result := scheduler.Run()
	// 输出结果
	for _, task := range result {
		fmt.Printf("任务ID:%d 任务名字:%s 执行完成,耗时:%v\n", task.ID, task.Name, task.Duration)
	}
}
func main() {
	//Goroutine()
	TaskScheduler()
}
