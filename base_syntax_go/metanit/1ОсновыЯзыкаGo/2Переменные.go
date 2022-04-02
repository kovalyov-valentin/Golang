package main // 1

import "fmt"

func main() {
	var hello string
	hello = "Hello World"
	fmt.Println(hello)
}

package main // 2

import "fmt"

func main() {
	var hello string = "Hello World"
	fmt.Println(hello)
}

package main // 3 

import "fmt"

func main() {
	var (
		name string = "Valentin"
		age int = 24
	)
	fmt.Println(name)
	fmt.Println(age)
}

package main // 4

import "fmt"

func main() {
	var hello string = "Hello World"
	fmt.Println(hello)
	hello = "Hello Go"
	fmt.Println(hello)
	hello = "Go Go Go Ole Ole Ole"
	fmt.Println(hello)

}

package main // 5

import "fmt"

func main() {
	name := "Valentin"
	fmt.Println(name)
}