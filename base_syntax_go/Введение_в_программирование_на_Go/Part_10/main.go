package main

import "fmt"
import m "github.com/Golang/base_syntax_go/Введение_в_программирование_на_Go/Part_10/math"

func main() {
	xs := []float64{1,2,3,4}
	avg := m.Average(xs)
	fmt.Println(avg)
}