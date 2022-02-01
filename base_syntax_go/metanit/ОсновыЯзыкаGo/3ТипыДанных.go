//Целочисленные типы


// package main

// import "fmt"

// func main() {
// 	var a int8 = -1
// 	var b uint8 = 2
// 	var c byte = 3  // byte - синоним типа uint8
// 	var d int16 = -4	
// 	var f uint16 = 5	
// 	var g int32 = -6
// 	var h rune = -7     // rune - синоним типа int32	
// 	var j uint32 = 8	
// 	var k int64 = -9
// 	var l uint64 = 10
// 	var m int = 102
// 	var n uint = 105
// 	fmt.Println(a, b, c, d, f, g, h, j, k, l, m, n)
// }


// Числа с плавающей точкой

// package main

// import "fmt"

// func main() {
// 	var f float32 = 18
// 	var g float32 = 4.5
// 	var d float64 = 0.23
// 	var pi float64 = 3.14
// 	var e float64 = 2.7
// 	fmt.Println(f, g, d, pi, e)
// }

//Комплексные числа

// package main

// import "fmt"

// func main() {
// 	var f complex64 = 1+2i
// 	var g complex128 = 4+3i
// 	fmt.Println(f, g)
// }

//Тип bool

// package main

// import "fmt"

// func main() {
// 	var isAlive bool = true
// 	var isEnabled bool = false
// 	fmt.Println(isAlive, isEnabled)
// }

//Строки

// package main

// import "fmt"

// func main() {
// 	var name string = "Том Сойер"
// 	fmt.Println(name)
// }

package main

import "fmt"

func main() {
	var (
		name = "Valentin"
		age = 24
	) 

	fmt.Println(name, age)
}