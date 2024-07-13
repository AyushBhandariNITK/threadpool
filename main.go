package main

import (
	"context"
	"math/rand"
	"threadpool/examples"
	"threadpool/executor"
	"threadpool/log"
)

func main() {
	log.Print(log.Info, "Program has started!!!")
	pool := executor.NewThreadPool()
	pool.Builder().SetWorkers(8).SetTasks(100)
	pool.Start()
	log.Print(log.Info, "pool %+v", pool)

	for i := 0; i < 100; i++ {
		go func() {
			row := rand.Intn(100) + 1
			col := rand.Intn(100) + 1
			task := examples.NewMatrix(row, col)
			pool.Submit(context.Background(), executor.NewJob(task.Execute))
		}()

	}
}
