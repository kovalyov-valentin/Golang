package main 

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}
func CircleArea(c *Circle) float64{
	return math.Pi * c.r*c.r
}
func main() {
	var c Circle = Circle {x:0, y:0, r:5}
	fmt.Println(CircleArea(&c))
}