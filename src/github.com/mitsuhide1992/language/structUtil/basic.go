package structUtil

import (
	// "fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func main() {
// 	var v = &Vertex{3, 4}
// 	fmt.Println(v.abs())
// }
