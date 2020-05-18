package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool管理一组可以安全在多个goroutine间共享的资源
type Pool struct {
	m	sync.Mutex
	resources	chan io.Closer
	factory	func() (io.Closer, error)
	closed	bool
}

// show that acquiring from an already closed pool
var ErrPoolClosed = errors.New("Pool has been closed")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Wrong size value")
	}

	return &Pool{
		factory: fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (this *Pool) Acquire() (io.Closer, error) {
	select {
		case r, ok := <-this.resources:
			log.Println("Acquire:", "Shared Resource")
			if !ok {
				return nil, ErrPoolClosed
			}
			return r, nil

		default:
			log.Println("Acquire:", "New Resource")
			return this.factory()
	}
}

func (this *Pool) Release(r io.Closer) {
	this.m.Lock()
	defer this.m.Unlock()

	if this.closed {
		r.Close()
		return
	}

	select {
		case this.resources <- r:
			log.Println("Release:", "In Queue")
		default:
			log.Println("Release:", "Closing")
			r.Close()
	}
}

func (this *Pool) Close() {
	this.m.Lock()
	defer this.m.Unlock()

	if this.closed {
		return
	}

	this.closed = true

	// 在清空通道资源之前，将通道关闭
	// 否则会发生死锁
	close(this.resources)

	for r := range this.resources {
		r.Close()
	}
}