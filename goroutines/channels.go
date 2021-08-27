package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for x:= 0; ; x++ {
		out <- x
		time.Sleep(time.Second * 1)
	}
	close(out)
}
func squarer(squares chan<- int, in <-chan int) {
	for v:= range in {
		squares <- v * v
	}
	close(squares)
}
func printer(in <- chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

/*func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// generation
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(time.Second * 1)
		}
	}()
	// squares
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()
	// printing
	for {
		fmt.Println(<- squares)
	}
}*/
