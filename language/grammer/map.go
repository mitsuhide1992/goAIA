package main

import "fmt"

func main() {
	var arr = [10]int{1, 2}
	fmt.Println(arr[0])

	var m = map[string]string{
		"google": "lolo",
		"baidu":  "polo",
	}
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

	var k int = 8
	fmt.Println(k)
}
