//Другая вещь, которую мы можем изменить в нашей программе - это цикл:

package main

import "fmt"

func main() {
	var x [5]float64 
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for i, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

//В этом цикле i представляет текущую позицию в массиве, а value будет тем же самым что и x[i]. 
//Мы использовали ключевое слово range перед переменной, по которой мы хотим пройтись циклом.
//Выполнение этой программы вызовет другую ошибку:

//$ go run tmp.go
//# command-line-arguments
//.\tmp.go:16: i declared and not used

//Компилятор Go не позволяет вам создавать переменные, которые никогда не используются в коде. 
//Поскольку мы не используем i внутри нашего цикла, то надо изменить код следующим образом: (Подробнее в test6)