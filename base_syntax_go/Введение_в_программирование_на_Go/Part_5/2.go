//Длина среза, созданного таким образом: make([]int, 3, 9) равняется трем.

package main

import "fmt"

func main() {
	slice := make([]int, 3, 9)
	fmt.Println(slice)
}


