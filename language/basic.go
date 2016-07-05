package main

import "fmt"
// import "math"

// go的基本类型
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 的别名

// rune // int32 的别名
//      // 代表一个Unicode码

// float32 float64

// complex64 complex128

const Pi = 3.14

func main() {
	var p, q int = 2, 3
	c := add(2, 3)
	fmt.Println(c)

	a := 3;
	b := 2;

	a, b = swap(a, b)
	fmt.Println(a, b, 3)
	fmt.Println(p, q)
	// fmt.Printf("Hello, world，ZLY,%g\n", math.Sqrt(16))
}


func add(a int, b int) int {
	return a + b
}

func swap (a, b int) (x int, y int) {
	x = b
	y = a
	return
	// return b, a
}