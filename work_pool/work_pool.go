package work_pool

import "sync"

type WorkPool interface {
	Submit(work func())
	Start()
	Stop()
}

type workPool struct {
	numWorkers int
	workQueue  chan func()
	wg         sync.WaitGroup
}

func NewWorkPool(numWorkers int) WorkPool {
	return &workPool{
		numWorkers: numWorkers,
		workQueue: make(chan func()),
	}
}

func (p *workPool) Submit(work func()) {
	p.wg.Add(1)
	p.workQueue <- work
}

func (p *workPool) Start() {
	for i := 0; i < p.numWorkers; i++ {
		go p.work()
	}
}

func (p *workPool) Stop() {
	close(p.workQueue)
	p.wg.Wait()
}

func (p *workPool) work() {
	for {
		select {
		case task, ok := <-p.workQueue:
			if ok {
				defer p.wg.Done()
				task()
			} else {
				return
			}
		}
	}
}