// package main

// import "fmt"

// func main() {
// 	var users []string
// 	fmt.Println(users)
// }

// package main

// import "fmt"

// func main() {
// 	var users = []string{"Tom,", "Alice,", "Kate"}
// 	users2 := []string{"Tom,", "Alice,", "Kate"}
// 	fmt.Println(users)
// 	fmt.Println(users2)
// }

// package main

// import "fmt"

// func main() {
// 	var users []string = []string{"Tom", "Alice", "Kate"}
// 	fmt.Println(users[2])
// 	users[2] = "Katherine"

// 	for _, value := range users {
// 		fmt.Println(value)
// 	}
// }

// package main

// import "fmt"
// func main() {
// 	var users []string = make([]string, 3)
// 	users[0] = "Tom"
// 	users[1] = "Alice"
// 	users[2] = "Bob"
// 	fmt.Println(users[0])
// 	fmt.Println(users[1])
// 	fmt.Println(users[2])
// }

// Добавление в срез

// package main

// import "fmt"

// func main() {
// 	users := []string{"Tom", "Alice", "Kate"}
// 	users = append(users, "Bob")

// 	for _, value := range users {
// 		fmt.Println(value)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Mike", "Paul","Robert"}
// 	users1 := initialUsers[2:6]
// 	users2 := initialUsers[:4]
// 	users3 := initialUsers[3:]

// 	fmt.Println(users1)
// 	fmt.Println(users2)
// 	fmt.Println(users3)
// }


package main

import "fmt"

func main() {
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Mike", "Paul","Robert"}
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users)
}