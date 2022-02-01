// package main

// import "fmt"

// func main() {
// 	var numbers [5] int
// 	fmt.Println(numbers)
// }

// package main

// import "fmt"

// func main() {
// 	var numbers [5]int = [5]int{1,2,3,4,5,}
// 	fmt.Println(numbers)
// }

// package main

// import "fmt"

// func main() {
// 	var numbers [5]int = [5]int{1,2}
// 	fmt.Println(numbers)
// }

// package main

// import "fmt"

// func main() {
// 	var numbers = [...]int{1,2,3,4,5}
// 	numbers2 := [...]int{1,2,3}
// 	fmt.Println(numbers)
// 	fmt.Println(numbers2)
// }

// package main

// import "fmt"

// func main() {
// 	var numbers [3]int = [3]int{1,2,3}
// 	var numbers2 [4]int = [4]int{1,2,3,4}
// 	numbers = numbers2 //В данном случае при присвоении мы получим ошибку, так как данные одного типа пытаемся передать переменной другого типа.
// 	fmt.Println(numbers)
// 	fmt.Println(numbers2)
// }

//Индексы

// package main

// import "fmt"

// func main() {
// 	var numbers [5]int = [5]int{1,2,3,4,5}
// 	fmt.Println(numbers[0])
// 	fmt.Println(numbers[4])
// 	numbers[0] = 87
// 	fmt.Println(numbers[0])
// }

package main

import "fmt"

func main() {
	var colors = [...]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2])
}