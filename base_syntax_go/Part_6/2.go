/*Напишите функцию, которая принимает число, делит его пополам и возвращает true в случае, если исходное число чётное, и false, если нечетное. 
Например, half(1) должна вернуть (0, false), в то время как half(2) вернет (1, true).*/

package main

import "fmt"

func main() {
	half(10)
}

func half(my int) {

	half_int := my / 2

	if half_int % 2 == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

