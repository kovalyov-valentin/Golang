// package main

// import ("fmt"; "math")

// func distance(x1, y1, x2, y2 float64) float64 {
//     a := x2 - x1
//     b := y2 - y1
//     return math.Sqrt(a*a + b*b)
// }
// func rectangleArea(x1, y1, x2, y2 float64) float64 {
//     l := distance(x1, y1, x1, y2)
//     w := distance(x1, y1, x2, y1)
//     return l * w
// }
// func circleArea(x, y, r float64) float64 {
//     return math.Pi * r*r
// }
// func main() {
//     var rx1, ry1 float64 = 0, 0
//     var rx2, ry2 float64 = 10, 10
//     var cx, cy, cr float64 = 0, 0, 5

//     fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
//     fmt.Println(circleArea(cx, cy, cr))
// }


// Добавьте новый метод perimeter в интерфейс Shape, который будет вычислять периметр фигуры. 
// Имплементируйте этот метод для Circle и Rectangle.

package main

import (
	"fmt" 
	"math"
)
type Shape interface{
	perimeter() float64
}
type Circle struct {
	x, y, r float64
}
type Restangle struct {
	x1, y1, x2, y2 float64
}
func totalPerimeter(shapes ...Shape) float64 {
    var perimeter float64
    for _, p := range shapes {
        perimeter += p.perimeter()
    }
    return perimeter
}
func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}
func restangleArea(r *Restangle) float64{
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (r *Restangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l + w + l + w
}
func circleArea(c *Circle) float64{
	return math.Pi * c.r*c.r
}
func (c *Circle) perimeter() float64{
	return 2 * math.Pi * c.r
}
func main() {
	var c Circle = Circle {x:0, y:0, r:5}
	fmt.Println(circleArea(&c))
	fmt.Println(totalPerimeter(&c))
	r := Restangle{0, 0, 10, 10}
	fmt.Println(restangleArea(&r))
	fmt.Println(totalPerimeter(&r))
	// c := new(Circle)
	// c := Circle {x: 0, y: 0, r: 5}
	// fmt.Println(c.x, c.y, c.r)
	// c.x = 10
	// c.y = 5

}