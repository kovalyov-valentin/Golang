//Карта (также известна как ассоциативный массив или словарь) — это неупорядоченная коллекция пар вида ключ-значение. Пример:

//var x map[string]int
//Карта представляется в связке с ключевым словом map, следующим за ним типом ключа в скобках и типом значения после скобок. 
//Читается это следующим образом: «x — это карта string-ов для int-ов».

//Подобно массивам и срезам, к элементам карт можно обратиться с помощью скобок. Запустим следующую программу:

//var x map[string]int
//x["key"] = 10
//fmt.Println(x)

//package main

//import "fmt"

//func main() {
//	var x map[string]int
//	x["key"] = 10
//	fmt.Println(x)
//}

//Вы должны увидеть ошибку, похожую на эту:
//panic: runtime error: assignment to entry in nil map

//goroutine 1 [running]:
//main.main()
//  main.go:7 +0x4d

//goroutine 2 [syscall]:
//created by runtime.main
//        C:/Users/ADMINI~1/AppData/Local/Temp/2/bindi
//t269497170/go/src/pkg/runtime/proc.c:221
//exit status 2

//До этого момента мы имели дело только с ошибками во время компиляции. Сейчас мы видим ошибку исполнения.
//Проблема нашей программы в том, что карта должна быть инициализирована перед тем, как будет использована. Надо написать так:

//package main

//import "fmt"

//func main() {
//	x:= make(map[string]int)
//	x["key"] = 10
//	fmt.Println(x["key"])
//}

//Если выполнить эту программу, то вы должны увидеть 10. Выражение x["key"] = 10 похоже на те, 
//что использовались при работе с массивами, но ключ тут не число, а строка (потому что в карте указан тип ключа string). 
//Мы также можем создать карты с ключом типа int:

//package main

//import "fmt"

//func main(){
//	x := make(map[int]int)
//	x[1] = 10
//	fmt.Println(x[1])
//}

//Это выглядит очень похоже на массив, но существует несколько различий. 
//Во-первых, длина карты (которую мы можем найти так: len(x)) может измениться, когда мы добавим в нее новый элемент. 
//В самом начале при создании длина 0, после x[1] = 10 она станет равна 1. Во-вторых, карта не является последовательностью. 
//В нашем примере у нас есть элемент x[1], в случае массива должен быть и первый элемент x[0], но в картах это не так.
//Также мы можем удалить элементы из карты используя встроенную функцию delete:

//delete(x, 1)