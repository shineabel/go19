package work

import "sync"

type Worker interface {
	Task()
}


type Pool struct {
	wg7 sync.WaitGroup
	worker chan Worker
}



func New(routineCount int) *Pool  {


	p := Pool{
		worker:make(chan Worker),
	}

	p.wg7.Add(routineCount)

	for i := 0; i < routineCount; i++{
		go func() {
			for w := range p.worker {
				w.Task()
			}
			p.wg7.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker)  {
	p.worker <- w
}

func (p *Pool) Shudown()  {
	close(p.worker)
	p.wg7.Wait()

}

func main() {

}