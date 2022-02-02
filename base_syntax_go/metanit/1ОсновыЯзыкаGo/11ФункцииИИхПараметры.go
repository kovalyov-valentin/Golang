// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello Go")
// }

// package main

// import "fmt"

// func main() {
// 	hello()
// 	hello()
// 	hello()

// }
// func hello() {
// 	fmt.Println("Hello World")
// }

// Параметры функции

// package main

// import "fmt"

// func main() {
// 	add(4,5)
// 	add(20,6)
// }
// func add(x int, y int) {
// 	var z = x + y
// 	fmt.Println("x + y = ", z)
// }

// package main

// import "fmt"

// func main() {
// 	add(1, 2, 3.4, 5.6, 1.2)
// }
// func add(x, y int, a, b, c float32) {
// 	var z = x + y
// 	var d = a + b + c
// 	fmt.Println("x + y = ", z)
// 	fmt.Println("a + b + c = ", d)
// }

// package main

// import "fmt"

// func main() {
// 	a := 8
// 	fmt.Println("a before: ", a)
// 	increment(a)
// 	fmt.Println("a after: ", a)
// }
// func increment(x int) {
// 	fmt.Println("x before: ", x)
// 	x += 20
// 	fmt.Println("x after: ", x)
// }

// Неопределенное количество параметров

// package main

// import "fmt"

// func main() {
// 	add(1,2,3)
// 	add(1,2,3,4)
// 	add(5,6,7,2,3)
// }
// func add(numbers ...int) {
// 	var sum = 0
// 	for _, number := range numbers {
// 		sum += number
// 	}
// 	fmt.Println("sum: ", sum)
// }