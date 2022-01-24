//Напишите программу, которая меняет местами два числа (x := 1; y := 2; swap(&x, &y) должно дать x=2 и y=1).

package main

import "fmt"

func main(){
	fmt.Println("x = 1, y = 2")
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println("x = ", x, "y = ", y)
}
func swap(x *int, y *int) {
	var temporary int
	temporary = *y
	*y = *x
	*x = temporary
}