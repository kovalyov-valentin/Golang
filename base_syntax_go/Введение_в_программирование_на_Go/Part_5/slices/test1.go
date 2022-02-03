//В Go есть две встроенные функции для срезов: append и copy. Вот пример работы функции append:
//func main() {
//    slice1 := []int{1,2,3}
//    slice2 := append(slice1, 4, 5)
//    fmt.Println(slice1, slice2)
//}

//После выполнения программы slice1 будет содержать [1,2,3], а slice2 — [1,2,3,4,5]. 
//append создает новый срез из уже существующего (первый аргумент) и добавляет к нему все следующие аргументы.

//Пример работы copy:
//func main() {
//    slice1 := []int{1,2,3}
//    slice2 := make([]int, 2)
//    copy(slice2, slice1)
//    fmt.Println(slice1, slice2)
//}
//После выполнения этой программы slice1 будет содержать [1,2,3], а slice2 — [1,2]. 
//Содержимое slice1 копируется в slice2, но поскольку в slice2 есть место только для двух элементов, то только два первых элемента slice1 будут скопированы.


//package main

//import "fmt"

//func main() {
//	slice1 := []int{1,2,3}
//	slice2 := append(slice1, 4, 5)
//	fmt.Println(slice1, slice2)  
//}


package main

import "fmt"

func main() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}