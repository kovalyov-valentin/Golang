// package main

// import "fmt"

// func main() {
// 	var a = add(4,5)
// 	var b = add(20, 6)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
// func add(x, y int) int {
// 	return x + y
// }

// Именованные возвращаемые результаты

// package main

// import "fmt"

// func main() {
// 	var a = add(4,5)
// 	var b = add(20,6)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
// func add (x, y int) (z int) {
// 	z = x + y
// 	return
// }

// package main

// import "fmt"

// func main() {
// 	var a = add(4,5)
// 	var b = add(20,6)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
// func add(x, y int) int {
// 	var z int = x + y
// 	return z
// }

// Возвращение нескольких значений

package main

import "fmt"

func main() {
	var age, name = add(20,4,"Valentin ", "Kovalev")
	fmt.Println(age)
	fmt.Println(name)
}
func add(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + "" + lastName
	return z, fullName
}