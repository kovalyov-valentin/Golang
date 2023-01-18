//Последовательность чисел Фибоначчи определяется как fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Напишите рекурсивную функцию, находящую fib(n).

package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}
func fib(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}
	
}