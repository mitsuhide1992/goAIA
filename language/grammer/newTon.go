package main 

import "fmt"
import "math"
/*
 * 牛顿开方迭代，golang官网的题
 */

func main() {
	fmt.Println(sqrt(12.0))
}

func sqrt(x float64) (y float64) {
	var z float64 = float64(1)

	var z2 float64 = newTon(z, x)

	for (math.Abs(z2 - z) > 0.00001) {
		z = z2
		z2 = newTon(z, x)
	}

	y = z2
	// y = x
	return
}

func newTon(a float64, x float64) (b float64) {
	b = a - (a * a - x)/(2 * a)
	return
}
