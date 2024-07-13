package executor

import (
	"sync"
	"threadpool/log"
)

type ThreadPool struct {
	UUID           string
	threadPoolName string
	workers        int
	jobs           int
	start          *sync.Once
	close          *sync.Once
}

func NewThreadPool() *ThreadPool {
	uuid := GenerateUUID()
	return &ThreadPool{
		UUID:           uuid,
		threadPoolName: GenerateName(uuid),
		start:          &sync.Once{},
		close:          &sync.Once{},
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

func (p *ThreadPool) SetJobs(jobs int) *ThreadPool {
	p.jobs = jobs
	return p
}

func (p *ThreadPool) Start() {

	p.start.Do(func() {
		log.Print(log.Info, "Threadpool with ID: %s started !!! ", p.threadPoolName)
	})

}

func (p *ThreadPool) Close() {
	p.close.Do(func() {
		log.Print(log.Info, "Threadpool with ID: %s closed !!!", p.threadPoolName)
	})
}
