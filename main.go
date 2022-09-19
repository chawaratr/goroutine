package main

import (
	"fmt"

	"github.com/chawaratr/goroutine/pkg/eroutine"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("hello")

	g := eroutine.New()
	g.Begin()
	for i := 0; i < 10; i++ {
		g.Do(func() {
			fmt.Println("inside cb eroutine")
			for i := 0; i < 10000000000; i++ {

			}
		})
		g = nil
	}

	g.End()

	fmt.Println(g)
}
