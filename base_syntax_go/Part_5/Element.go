//Напишите программу, которая находит самый наименьший элемент в этом списке:

//package main

//import "fmt"

//func main() {
//	x := []int{ 
//		48, 96, 86, 68,
//		57, 82, 63, 70,
//		37, 34, 83, 27,
//		19, 97, 9, 17,
//	}
//	var total int = 0
//	for i := 0; i < len(x); i ++{
//		total += x[i]
//	}
//	fmt.Println(total / len(x))
	


//package main

//import "fmt"

//func main() {
//	x := []int{
//		48, 96, 86, 68,
//		57, 82, 63, 70,
//		37, 34, 83, 27,
//		19, 97, 9, 17,
//	}
//	fmt.Println(x[len(x)-1])
//}

package main

import "fmt"

func main() {
	var x = []int{

	48, 96, 86, 68,
	57, 82, 63, 70,
	37, 34, 83, 27,
	19, 97, 9, 17,
	}

	min := x[1]
	for _, value := range x {

		if value < min {
			min = value
		}
	}

	fmt.Println(min)
}