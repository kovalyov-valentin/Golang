//Напишите функцию с переменным числом параметров, которая находит наибольшее число в списке.

package main

import "fmt"

func main(){
	task(1,2,3,4,5,6,7,8,9,10)
}
func task(args ...int) {
	max := 0

	for _, value := range args {
		if value > max {
			max = value
		}
	}

	fmt.Println("Maximum is: ", max)
}