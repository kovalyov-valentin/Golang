// package main

// import "fmt"

// func main() {
// 	var x int = 4
// 	var p *int
// 	p = &x
// 	fmt.Println(p)
// }

// package main

// import "fmt"

// func main() {
// 	var x int = 4
// 	var p *int = &x
// 	fmt.Println("Adress:", p)
// 	fmt.Println("Value:", *p)
// }

// package main

// import "fmt"

// func main() {
// 	var x int = 4
// 	var p *int = &x
// 	*p = 25
// 	fmt.Println(x)
// }

// package main

// import "fmt"

// func main() {
// 	f := 2.3
// 	pf := &f
// 	fmt.Println("Adress:", pf)
// 	fmt.Println("Value:", *pf)
// }

// Пустой указатель

// package main

// import "fmt"

// func main() {
// 	var pf *float64
// 	fmt.Println("Value:", *pf) // ! ошибка, указатель не указывает на какой-либо объект
// }

// package main

// import "fmt"

// func main() {
// 	var x float64 = 2.5
// 	var pf *float64 = &x
// 	if pf != nil{
// 		defer fmt.Println("Value:", *pf)
// 	}
// 	fmt.Println("Adress:", pf)
// }

// Функция new

// package main

// import "fmt"

// func main() {
// 	p := new(int)
// 	fmt.Println("Value:", *p)
// 	*p = 8
// 	fmt.Println("Value:", *p)
// }