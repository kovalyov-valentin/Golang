// package main 

// import (
// 	"os"
// 	"strconv"
// 	"log"
// )

// func main() {
// 	sum := 0
// 	for i, arg := range os.Args[1:] {
// 	  n, err := strconv.Atoi(arg)
// 	  if err != nil {
// 		log.Fatalf("can't parse arg #%d '%s'", i, arg)
// 	  }
  
// 	  sum += n
// 	}
  
// 	log.Printf("sum is %d", sum)
//   }

package main

import (
	"fmt"
)

func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	BubbleSort(ar)
	fmt.Println(ar)
}

func BubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				swap(ar, i, j)
			}
		}
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}