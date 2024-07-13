package main

import (
	"threadpool/executor"
	"threadpool/log"
)

func main() {
	log.Print(log.Info, "Program has started!!!")
	x := executor.NewThreadPool()
	x.Builder().SetWorkers(4).SetJobs(4)
	x.Start()
	log.Print(log.Info, "pool %+v", x)
}
