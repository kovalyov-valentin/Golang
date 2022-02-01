// package main

// import "fmt"

// func main() {
// 	defer finish()
// 	fmt.Println("Program has been starded")
// 	fmt.Println("Program is working")
// }
// func finish() {
// 	fmt.Println("Program has been finished")
// }

// package main

// import "fmt"

// func main() {
// 	defer finish()
// 	defer fmt.Println("Program has been started")
// 	fmt.Println("Program is working")
// }
// func finish() {
// 	fmt.Println("Program has been finished")
// }

// Panic

package main

import "fmt"

func main() {
	fmt.Println(divide(15, 5))
	fmt.Println(divide(4, 0))
	fmt.Println("Program has been finished")
}
func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}