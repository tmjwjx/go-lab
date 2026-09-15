package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Id int
}
type Scheduler []*Task

func (t *Task) Run(c context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("执行 %d 任务\n", t.Id)
	for i := 0; i < 10; i++ {
		select {
		case <-c.Done():
			fmt.Printf("任务 %d 已取消\n", t.Id)
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("任务 %d 正在执行...(%d/10)\n", t.Id, i+1)
		}
	}
	fmt.Printf("任务 %d 完成\n", t.Id)
}
func (s *Scheduler) RunTasks(c context.Context) {
	var w sync.WaitGroup
	for _, task := range *s {
		w.Add(1)
		go task.Run(c, &w)
	}
	w.Wait()
}

func main() {
	scheduler := Scheduler{}
	
	for i := 0; i < 5; i++ {
		scheduler = append(scheduler, &Task{i})
	}
	c, cancel := context.WithCancel(context.Background())
	go func() {
		
		time.Sleep(5 * time.Second)
		fmt.Println("取消所有任务")
		cancel()
	}()
	
	scheduler.RunTasks(c)
	fmt.Println("所有任务已完成或已取消")
}