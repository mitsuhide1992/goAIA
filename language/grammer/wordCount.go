package main

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

func main() {
	var str string = "I am in baidu am am"
	var re map[string]int = WordCount(str)
	fmt.Println(re)
}

func WordCount(s string) map[string]int {
	var strs [] string = strings.Fields(s)
	var result map[string]int = make(map[string]int)
	for _,v := range strs {
		_,ok := result[v]
		if ok == true {
			result[v]++
		} else {
			result[v]=1
		}
	}
	return result
}