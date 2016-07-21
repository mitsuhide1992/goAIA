package grammar

// import "fmt"

// go语言不用分号
// func main() {
// 	// var hy = func(x, y int) int {
// 	// 	return (x*x + y*y)
// 	// }
// 	// fmt.Println(compute(hy))

// 	var pos, neg func(int) int = closure(), closure()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(
// 			pos(i),
// 			neg(-2*i),
// 		)
// 	}
// }

//函数作为参数
func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

//闭包
func closure() func(int) int {
	var sum int = 0
	return func(x int) int {
		sum += x
		return sum
	}
}
