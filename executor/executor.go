package executor

import (
	"context"
	"sync"
	"threadpool/log"
	"time"
)

type ThreadPool struct {
	UUID           string
	threadPoolName string
	workers        int
	tasks          int
	start          *sync.Once
	close          *sync.Once
	quit           chan struct{}
	TaskChan       chan Task
}

func NewThreadPool() *ThreadPool {
	uuid := GenerateUUID()
	return &ThreadPool{
		UUID:           uuid,
		threadPoolName: GenerateName(uuid),
		start:          &sync.Once{},
		close:          &sync.Once{},
		quit:           make(chan struct{}),
	}
}

func (p *ThreadPool) Builder() *ThreadPool {
	return p
}

func (p *ThreadPool) SetThreadPoolId(name string) {
	UUIDNameMap.Store(p.UUID, name)
	p.threadPoolName = name
}
func (p *ThreadPool) SetWorkers(workers int) *ThreadPool {
	p.workers = workers
	return p
}

func (p *ThreadPool) SetTasks(tasks int) *ThreadPool {
	p.tasks = tasks
	p.TaskChan = make(chan Task, tasks)
	return p
}

func (p *ThreadPool) Start() {
	p.start.Do(func() {
		log.Print(log.Info, "Threadpool with ID: %s started !!! ", p.threadPoolName)
		p.StartWorkers()
	})
}

func (p *ThreadPool) Close() {
	log.Print(log.Info, "Stopping all workers !!!")
	p.close.Do(func() {
		close(p.quit)
		log.Print(log.Info, "Threadpool with ID: %s closed !!!", p.threadPoolName)
	})
}

func (p *ThreadPool) StartWorkers() {
	for i := 0; i < p.workers; i++ {
		go func(workerId int) {
			log.Print(log.Info, "Worker %d started", workerId)
			for {
				select {
				case <-p.quit:
					log.Print(log.Info, "Closing worker %d due to quit chan closed", workerId)
					return
				case job := <-p.TaskChan:
					err := job.Execute(context.Background())
					if err != nil {
						log.Print(log.Info, "For Task %s , Error occurred: %s ", job.GetId(), err.Error())
					}
				}
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func (p *ThreadPool) Submit(ctx context.Context, task Task) {
	select {
	case p.TaskChan <- task:
	case <-p.quit:
	}
}
