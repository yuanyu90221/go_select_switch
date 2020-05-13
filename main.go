package main

import "fmt"

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
