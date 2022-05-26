package main 

import (
	"fmt"
	"sort"
)

func main() {
	arrays := []int{2, 1, 3, 3, 2, 1, 1, 2}
	arr := len(arrays)
	fmt.Println(arr)
	sort.Ints(arrays)
	fmt.Println(arrays)
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arrays := []int{2, 1, 3, 3, 2, 1}
// 	arr := len(arrays)
// 	fmt.Println(arr)
// }