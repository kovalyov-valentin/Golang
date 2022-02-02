// package main

// import "fmt"

// type contact struct {
// 	email string
// 	phone string
// }
// type person struct {
// 	name string
// 	age int
// 	contactInfo contact
// }
// func main() {
// 	var tom = person {
// 		name: "Tom",
// 		age: 24,
// 		contactInfo: contact {
// 			email: "tom@gmail.com",
// 			phone: "+1234567899",
// 		},
// 	}
// 	tom.contactInfo.email = "supertom@gmail.com"

// 	fmt.Println(tom.contactInfo.email)
// 	fmt.Println(tom.contactInfo.phone)
// }

// package main

// import "fmt"

// type contact struct {
// 	email string
// 	phone string
// }
// type person struct {
// 	name string
// 	age int
// 	contact
// }
// func main() {
// 	var tom = person {
// 		name: "Tom",
// 		age: 24,
// 		contact: contact {
// 			email: "tom@gmail.com",
// 			phone: "+1234567899",
// 		},
// 	}
// 	tom.email = "supertom@gmail.com"

// 	fmt.Println(tom.email)
// 	fmt.Println(tom.phone)
// }

// Хранения ссылки на структуру того же типа

package main

import "fmt"

type node struct {
	value int
	next *node
}
func printNodeValue(n *node) {
	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}
func main() {
	first := node {value: 4}
	second := node {value: 5}
	third := node {value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}