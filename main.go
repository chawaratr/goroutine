package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/chawaratr/goroutine/pkg/eroutine"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("hello")

	//Start Routine Here
	gRoutine := eroutine.New()

	var err error
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 10)
		err = gRoutine.Try(func() error {
			fmt.Println("inside cb eroutine")
			time.Sleep(time.Second * 5)

			return errors.New("Business error")
		})
		if err != nil {
			buildError(err)
			return
		}
	}

	err = gRoutine.End()
	if err != nil {
		buildError(err)
		return
	}
	//!End Routine Here

	fmt.Println("Normal End")

}

func buildError(err error) {
	fmt.Println("Main process with error:", err.Error())
}
