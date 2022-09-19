package eroutine

import (
	"fmt"
	"sync"
)

type eroutine struct {
	wg      sync.WaitGroup
	cerror  chan error
	counter int
}

type Igoroutine interface {
	Begin()
	Do(cb func())
	End()
}

func New() Igoroutine {
	x := new(eroutine)
	return x
}

func (g *eroutine) Begin() {
	// g.wg = sync.WaitGroup{}
	g.cerror = make(chan error, 1)
}

func (g *eroutine) Do(cb func()) {
	g.counter++
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		cb()
	}()
}

func (g *eroutine) End() {
	g.wg.Wait()
	fmt.Println("Count=", g.counter)
}
