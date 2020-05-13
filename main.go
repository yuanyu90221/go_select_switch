package main

import (
	"fmt"
	"time"
)

func main() {
	i = 100
	convert(i)
	i = float64(45.55)
	convert(i)
	i = "foo"
	convert(i)
	convert(float32(10.0))

	// channel
	ch := make(chan int, 1)

	// ch <- 1
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("exit")
	}
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	ch1 := make(chan int)
	select {
	case <-ch1:
	case <-timeout:
		fmt.Println("timeout 01")
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 02")
	}
}

var (
	i interface{}
)

func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is interger", t)
	case string:
		println("i is string", t)
	case float64:
		println("i is float64", t)
	default:
		println("type not found")
	}
}
