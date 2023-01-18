package main 

import (
	"fmt"
)

func main() {
	array := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(array)
	fmt.Println(array[3][2])
}