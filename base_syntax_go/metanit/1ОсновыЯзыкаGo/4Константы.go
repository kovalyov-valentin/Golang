// package main

// import "fmt"

// func main() {
// 	const pi float64 = 3.1415
// 	fmt.Println(pi)
// }

// package main

// import "fmt"

// func main() {
// 	const pi float64 = 3.1415
// 	pi = 2.7182 // Здесь будте ошибка, так как мы не можем изменить значение константы
// 	fmt.Println(pi)
// }

// package main

// import "fmt"

// func main() {
// 	const (
// 		pi = 3.1415
// 		e = 2.7182
// 	)
// 	fmt.Println(pi, e)
// }

// package main

// import "fmt"

// func main() {
// 	const (
// 		a = 1
// 		b 
// 		c
// 		d = 3
// 		f
// 	)
// 	fmt.Println(a, b, c, d, f)
// }

package main

import "fmt"

func main() {
	var j int = 7
	// const k = j  // Константа не может принимать значение переменной
	const s = 5
	const d = s
	fmt.Println(j, s, d)
}