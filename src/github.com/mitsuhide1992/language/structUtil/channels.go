// package main

// import "fmt"

// func sum(a []int, c chan int) {
// 	sum := 0
// 	for _, v := range a {
// 		sum += v
// 	}
// 	c <- sum // 将和送入 c
// }

// func main() {
// 	a := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(a[:len(a)/2], c)
// 	go sum(a[len(a)/2:], c)
// 	x, y := <-c, <-c // 从 c 中获取

// 	fmt.Println(x, y, x+y)
// }

package structUtil

// import "fmt"

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }
