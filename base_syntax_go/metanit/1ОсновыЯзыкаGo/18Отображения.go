// package main

// import "fmt"

// func main() {
// 	var people map[string]int
// 	fmt.Println(people)
// }

// package main

// import "fmt"

// func main() {
// 	var people = map[string]int {
// 		"Tom": 1,
// 		"Bob": 2,
// 		"Sam": 4,
// 		"Alice": 8,
// 	}
// 	fmt.Println(people)
// }

// package main

// import "fmt"

// func main() {
// 	var people = map[string]int {
// 		"Tom": 1,
// 		"Bob": 2,
// 		"Sam": 4,
// 		"Alice":8,
// 	}
// 	fmt.Println(people["Alice"])
// 	fmt.Println(people["Bob"])
// 	people["Bob"] = 32
// 	fmt.Println(people["Bob"])
// }

// package main

// import "fmt"

// func main() {
// 	var people = map[string]int {
// 		"Tom": 1,
// 		"Bob": 2,
// 		"Sam": 4,
// 		"Alice": 8,
// 	}
// 	if val, ok := people["Tom"]; ok {
// 		fmt.Println(val)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var people = map[string]int {
// 		"Tom": 1,
// 		"Bob": 2,
// 		"Sam": 4,
// 		"Alice": 8,
// 	}
// 	for key, value := range people{
// 		fmt.Println(key, value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	people := make(map[string]int)
// 	fmt.Println(people)
// }

// Добавление и удаление элементов

// package main

// import "fmt"

// func main() {
// 	var people = map[string]int {"Tom": 1, "Bob": 2}
// 	people["Kate"] = 128
// 	fmt.Println(people)
// }

package main

import "fmt"

func main() {
	var people = map[string]int {"Tom": 1, "Bob": 2, "Sam": 8}
	delete(people, "Bob")
	fmt.Println(people)
}