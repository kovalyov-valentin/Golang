// package main

// import "fmt"

// type library []string

// func(l library) print() {
// 	for _, val := range l {
// 		fmt.Println(val)
// 	}
// }
// func main() {
// 	var lib library = library{"Book1", "Book2", "Book3"}
// 	lib.print()
// }

package main

import "fmt"

type person struct {
	name string
	age int 
}
func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}
func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}
func main() {
	var tom = person{name: "Tom", age: 24}
	tom.print()
	tom.eat("борщ с капустой, но не красный")
}