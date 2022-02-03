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
	for i := 0; i <= len(x); i ++ {
		total += x[i]
	}
	fmt.Println(total / len(x))
}

//Нам выдаст ошибку, так как
//проблема в том, что len(x) и total имеют разный тип. total имеет тип float64, а len(x) — int. Так что, нам надо конвертировать len(x) в float64. (Подробнее в test4)