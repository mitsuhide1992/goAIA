package grammar

// import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
// fibonacci 只会调用一次，返回的是闭包函数，之后每次调用f都是调用闭包函数
func fibonacci() func() int {
	var a int = 1
	var b int = 1
	var c int = 0
	var i int = 0

	return func() int {
		i++
		if i == 0 {
			return 1
		} else if i == 1 {
			return 1
		} else {
			c = a + b
			a = b
			b = c
			return c
		}
	}
}

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
