// Мы скопировали функцию Average из главы 7 в наш новый пакет. Создайте Min и Max функции для нахождения наименьших и наибольших значений в срезах дробных чисел типа float64.

package main

import "fmt"

func main() {
    c := Min
    fmt.Println(c)
    b := Max
    fmt.Println(b)
}

func average(xs []float64) float64 {    
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}
func Max(xs []float64) float64 {
	max := xs[0]
	for _, val := range xs[1:] {
		if val > max {
            max := val
        }
	}
	return max
}
func Min(xs []float64) float64 {
	min := xs[0]
	for _, val := range xs[1:] {
		if val < min {
            min := val
        }
	}
	return min
}
