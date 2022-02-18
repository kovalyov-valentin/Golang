// В дополнение к спискам и картам, Го предоставляет еще несколько видов коллекций, доступных в пакете container. 
// В качестве примера рассмотрим container/list.

// СПИСОК
// Пакет container/list реализует двусвязный список.
// Каждый узел списка содержит значение и указатель на следующий узел.
// Но так как это двусвязный список, узел также содержит указатель на предыдущий.
// Такой список может быть создан с помощью следующей программы:

// package main

// import (
// 	"fmt"
// 	"container/list"
// )
// func main() {
// 	var x list.List 
// 	x.PushBack(1)
// 	x.PushBack(2)
// 	x.PushBack(3)

// 	for e := x.Front(); e != nil; e=e.Next() {
// 		fmt.Println(e.Value.(int))
// 	}
// }

// СОРТИРОВКА
// Пакет sort содержит функции для сортировки произвольных данных.
// Есть несколько предопределенных функций(для срезов, целочисленных значение и чисел с плавающей точкой).
// Вот пример, как отсортировать наши данные:

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}
type ByName []Person

func (this ByName) Len() int {
	return len(this) 
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func main() {
	kids := []Person {
		{"Jill",9},
		{"Jack",10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}