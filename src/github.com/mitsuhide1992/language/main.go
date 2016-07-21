package main

import (
	"fmt"
	"github.com/mitsuhide1992/language/structUtil"
)

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	structUtil.Fibonacci(c, quit)
}
