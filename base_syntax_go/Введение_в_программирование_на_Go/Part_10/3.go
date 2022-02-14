// Что такое псевдоним пакета и как его сделать?

Псевдоним пакеты позволяет ссылать на пакет с более коротким.

Сделать можно следующим образом:
Было:
package main

import "fmt"
import "github.com/Golang/base_syntax_go/Введение_в_программирование_на_Go/Part_10/math"

func main() {
    xs := []float64{1,2,3,4}
    avg := math.Average(xs)
    fmt.Println(avg)
}
Стало:
import m "github.com/Golang/base_syntax_go/Введение_в_программирование_на_Go/Part_10/math"

func main() {
	xs := []float64{1,2,3,4}
	avg := m.Average(xs)
	fmt.Println(avg)
}
Псевдоним указыывается без кавычек после слова import перед именем пакета: import m "math"