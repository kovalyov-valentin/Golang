package main 

import (
	"fmt"
)

// Определяется структура twoInts, имеющая два поля
type twoInts struct {
	X int64
	Y int64
}

// Определяем новую функцию regularFunction
// Она принимает два параметра типа twoInts
// Возвращает одно значение типа twoInts
func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

// Фукнция method() эквивалентна функции regularFunction()
// Однако method() - это метод типа
func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

// Как можно заметить способ вызова метода типа ( i.method(j) )
// отличается от способа вызова обычной функции (regularFunction(i, j))
func main() {
	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: -5, Y: -2}
	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
}
// Следует обратить внимание, что методы типа связаны с интерфейсами.
