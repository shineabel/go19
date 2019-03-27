package pool

import (
	"sync"
	"io"
	"errors"
	"log"
)

type Pool struct {
	closed bool
	m sync.Mutex
	resources chan io.Closer
	factory func()(io.Closer,error)
}

var ErrorPoolClosed = errors.New("Pool closed already")

func New(fn func()(closer io.Closer, err error), size uint)(*Pool,error)  {

	if(size <= 1){
		return nil,errors.New("size is too small")
	}
	return &Pool{
		factory:fn,
		resources:make(chan io.Closer,size),
	},nil
}

func (p *Pool) Acquire()(io.Closer,error)  {

	select {

	case r, ok := <- p.resources:
		log.Printf("shared")
		if(!ok){
			return nil,ErrorPoolClosed
		}
		return r,nil
	default:
		log.Printf("new resource")

		return p.factory()
	}
}

func (p *Pool)Release(r io.Closer)  {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed{
		r.Close()
		return
	}
	select {
	 case p.resources <- r:
	 	log.Printf("in queue")
	default:
		log.Printf("closing")
	 r.Close()
	}
}

func (p *Pool)Close()  {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed{
		return
	}
	p.closed = true
	close(p.resources)

	for r := range p.resources{
		r.Close()
	}
	
}

