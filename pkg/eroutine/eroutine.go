package eroutine

import (
	"fmt"
	"sync"
)

type eroutine struct {
	wg      sync.WaitGroup
	cerror  chan error
	counter int
	m       sync.Mutex
	quit    error
}

type Igoroutine interface {
	Try(cb func() error) error
	End() error
}

func New() Igoroutine {
	x := new(eroutine)
	x.cerror = make(chan error, 1)
	go x.catchError()
	return x
}

func (g *eroutine) Try(cb func() error) error {
	if g.quit != nil {
		return g.quit
	}
	counter := g.getCounter()
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		fmt.Println("inside eroutine no=", counter)

		err := cb()
		if err != nil {
			if len(g.cerror) < cap(g.cerror) {
				g.cerror <- err
			}
			return
		}
	}()
	return nil
}

func (g *eroutine) End() error {
	g.wg.Wait()
	return g.quit
}

func (g *eroutine) catchError() {
	for {
		select {
		case e := <-g.cerror:
			g.quit = e
		}
	}

}

func (g *eroutine) getCounter() int {
	g.m.Lock()
	g.counter++
	r := g.counter
	g.m.Unlock()
	return r
}
