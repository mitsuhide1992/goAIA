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

type Vertex struct {
	x int
	y int
}

func main() {
	var p, q int = 2, 3
	c := add(2, 3)
	fmt.Println(c)

	a := 3;
	b := 2;

	// defer fmt.Println("world")
	
	// 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }

	a, b = swap(a, b)
	fmt.Println(a, b, 3)
	fmt.Println(p, q)
	fmt.Println(Vertex{1,2})

	var vnode Vertex = Vertex{4,5}
	fmt.Println(vnode.x)
	fmt.Println(vnode.y)

	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{x: 1}  // Y:0 被省略
		v3 = Vertex{}      // X:0 和 Y:0
		poi  = &Vertex{1, 2} // 类型为 *Vertex
	)

	var pointer * Vertex = &vnode;
	fmt.Println("This is pointer",pointer.x)
	fmt.Println(v1,v2,v3,poi)
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