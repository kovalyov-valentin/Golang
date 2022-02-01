//Операции отношения

// package main

// import "fmt"

// func main() {
// 	var a int = 8
// 	var b int = 3
// 	var c bool = a == b
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 8
// 	var b int = 3
// 	var c bool = a > b
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 8
// 	var b int = 3
// 	var c bool = a < b
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 8
// 	var b int = 3
// 	var c bool = a <= b
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 8
// 	var b int = 3
// 	var c bool = a >= b
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 8
// 	var b int = 3
// 	var c bool = a != b
// 	var d bool = a != 8
// 	fmt.Println(c)
// 	fmt.Println(d)
// }

//Логические операции

// package main

// import "fmt"

// func main() {
// 	var a bool = true
// 	var b bool = !a
// 	var c bool = !b
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// package main

// import "fmt"

// func main() {
// 	var a bool = 4 > 5 && 6 > 8
// 	var b bool = 3 <= 5 && 10 > 8
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

package main

import "fmt"

func main() {
	var a bool = 4 > 5 || 6 > 8
	var b bool= 3 == 5 || 10 > 8
	fmt.Println(a)
	fmt.Println(b)
}