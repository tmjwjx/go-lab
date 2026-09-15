package main

import (
	"fmt"
)

type TaskScheduler struct {
	taskChannel chan func()
}

func (taskScheduler *TaskScheduler) AddTask(task func()) {
	taskScheduler.taskChannel <- task
}

func (taskScheduler *TaskScheduler) Start() {
	defer close(taskScheduler.taskChannel)

	for task := range taskScheduler.taskChannel {
		task()
	}
}

func main() {
	taskScheduler := TaskScheduler{
		taskChannel: make(chan func()),
	}

	task1 := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}
	task2 := func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%c\n", rune('A'+i))
		}
	}
	task3 := func() {
		fmt.Println("finish")
	}

	go taskScheduler.Start()

	taskScheduler.AddTask(task1)
	taskScheduler.AddTask(task2)
	taskScheduler.AddTask(task3)
}