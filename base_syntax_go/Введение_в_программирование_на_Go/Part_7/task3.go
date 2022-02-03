//Какое будет значение у переменной x после выполнения программы:

package main 

import "fmt"

func main() {
    one()
}
func one() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}
func square(x *float64) {
    *x = *x * *x
}

//Ответ: 2.25