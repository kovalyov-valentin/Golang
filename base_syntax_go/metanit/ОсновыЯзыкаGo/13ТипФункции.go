// package main

// import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }
// func main() {

// 	var f func(int,int) int = add
// 	fmt.Println(f(3, 4))

// 	var x = f(4, 5)
// 	fmt.Println(x)
// }

// package main

// import "fmt"

// func add(x int, y int) int{return x + y}
// func multiply(x int, y int) int{return x * y}

// func display(message string) {
// 	fmt.Println(message)
// }
// func main() {
// 	f := add
// 	fmt.Println(f(3,4))
// 	f = multiply
// 	fmt.Println(f(3,4))
// 	// f = display
// 	var f1 func(string) = display
// 	f1("hello")
// }

// Функции как параметры других функций

// package main

// import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }
// func multiply(x int, y int) int {
// 	return x * y
// }
// func action(n1 int, n2 int, operation func(int, int) int){
// 	result := operation(n1, n2)
// 	fmt.Println(result)
// }
// func main() {
// 	action(10, 25, add)
// 	action(5, 6, multiply)
// }

// package main

// import "fmt"

// func isEven(n int) bool {
// 	return n % 2 == 0
// }
// func isPositive(n int) bool {
// 	return n > 0
// }
// func sum(numbers []int, criteria func(int) bool) int {
// 	result := 0
// 	for _, val := range numbers {
// 		if(criteria(val)) {
// 			result += val
// 		}
// 	}
// 	return result
// }
// func main() {
// 	slice := []int{-2,4,3,-1,7,-4,23}
// 	sumOfEvens := sum(slice, isEven)
// 	fmt.Println(sumOfEvens)
// 	sumOfPositives := sum(slice, isPositive)
// 	fmt.Println(sumOfPositives)
// }

// Функция как результат другой функции

package main

import "fmt"

func add(x int, y int) int{return x + y}
func subtract(x int, y int) int{return x - y}
func multiply(x int, y int) int{return x * y}

func selectFn(n int) (func(int, int) int) {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}
func main() {
	f := selectFn(1)
	fmt.Println(f(3,4))
	f = selectFn(3)
	fmt.Println(f(3,4))
}