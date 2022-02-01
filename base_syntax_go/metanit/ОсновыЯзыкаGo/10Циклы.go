// package main

// import "fmt"

// func main() {
// 	for i := 1; i < 10; i ++ {
// 		fmt.Println(i * i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var i = 1
// 	for ; i < 10; i ++ {
// 		fmt.Println(i *i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var i = 1
// 	for ; i < 10;{
// 		fmt.Println(i * i)
// 		i++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var i = 1
// 	for i < 10{
// 		fmt.Println(i * i)
// 		i++
// 	}
// }

// Вложенные циклы

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 10; i ++ {
// 		for j := 1; j <= 10; j ++ {
// 			fmt.Print(i * j, "\t")
// 		}
// 		fmt.Println()
// 	}
// }

//Перебор массивов

// package main

// import "fmt"

// func main() {
// 	var users = [3]string{"Tom", "Alice", "Kate"}
// 	for index, value := range users{
// 		fmt.Println(index, value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var users = [3]string{"Tom", "Alice", "Kate"}
// 	for _, value := range users{
// 		fmt.Println(value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var users = [3]string{"Tom", "Alice", "Kate"}
// 	for i := 0; i < len(users); i++ {
// 		fmt.Println(users[2])
// 	}
// }

//Операторы break и continue

// package main

// import "fmt"

// func main() {
// 	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
// 	var sum = 0
// 	for _, value := range numbers {
// 		if value < 0 {
// 			continue
// 		}
// 		sum += value
// 	}
// 	fmt.Println("Sum:", sum)
// }

package main

import "fmt"

func main() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum = 0
	for _, value := range numbers {
		if value > 4 {
			break
		}
		sum += value
	}
	fmt.Println("Sum:", sum)
}