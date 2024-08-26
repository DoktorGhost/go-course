package main

import (
	"fmt"
	"sync"
)

type Task struct {
	name string
	work func()
}

func WorkerPool(n int, chanTask chan Task, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		go func(id int) {
			defer wg.Done()
			for task := range chanTask {
				fmt.Println("Worker ", id, " start")
				task.work()
				fmt.Println("Worker", id, "finished")
			}
		}(i)
	}
}

func main() {
	taskChan := make(chan Task)
	var wg sync.WaitGroup

	countWorker := 3
	wg.Add(countWorker)
	go WorkerPool(countWorker, taskChan, &wg)

	for i := 0; i < 10; i++ {
		i := i
		name := fmt.Sprintf("task#%d", i)
		funcTask := func() {
			for j := (i + 1) * 2; j < ((i+1)*2)+10; j++ {
				fmt.Println("Task#", i, " value", j)
			}
		}
		taskOne := Task{
			name: name,
			work: funcTask,
		}

		taskChan <- taskOne
	}
	close(taskChan)
	wg.Wait()
	fmt.Println("Finish")
}
