// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int

// Создание и инициализация структуры

// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int
// }
// func main() {
// 	var tom person = person{"Tom", 23}
// 	fmt.Println(tom)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int
// }
// func main() {
// 	var alice person = person{age: 23, name: "Alice"}
// 	fmt.Println(alice)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int
// }
// func main() {
// 	var tom = person {name: "Tom", age: 24}
// 	bob := person {name: "Bob", age: 31}
// 	fmt.Println(tom)
// 	fmt.Println(bob)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int
// }
// func main() {
// 	undefined := person{}
// 	fmt.Println(undefined)
// }

// Обращение к полям структуры

// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int
// }
// func main() {
// 	var tom = person {name: "Tom", age: 24}
// 	fmt.Println(tom.name)
// 	fmt.Println(tom.age)

// 	tom.age = 38
// 	fmt.Println(tom.name, tom.age)
// }

// Указатели на структуры

// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int
// }
// func main() {
// 	tom := person {name: "Tom", age: 22}
// 	// fmt.Println(tom.name, tom.age)
// 	var tomPointer *person = &tom
// 	tomPointer.age = 29
// 	fmt.Println(tom.age)
// 	(*tomPointer).age = 32
// 	fmt.Println(tom.age)
// }

// package main

// import "fmt"

// type person struct {
// 	name string
// 	age int
// }
// func main() {
// 	var tom *person= &person {name: "Tom", age: 23}
// 	var bob *person = new(person)
// 	fmt.Println(tom)
// 	fmt.Println(bob)
// }

package main

import "fmt"

type person struct {
	name string
	age int 
}
func main() {
	tom := person {name: "Tom", age: 22}
	var agePointer *int = &tom.age
	*agePointer = 35
	fmt.Println(tom.age)
}