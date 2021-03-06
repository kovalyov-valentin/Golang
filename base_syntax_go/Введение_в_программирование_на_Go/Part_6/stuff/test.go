/*package main

import "fmt"

func average(xs []float64) float64 {
	panic("Not Implemented")
}

func main() {
	xs := []float64{98,93,77,82,83}
	
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}


Функция average должна взять срез из нескольких float64 и вернуть один float64. Напишем перед функцией main:

func average(xs []float64) float64 {
    panic("Not Implemented")
}

Функция начинается с ключевого слова func, за которым следует имя функции. Аргументы (входы) определяются так: имя тип, имя тип, …. 
Наша функция имеет один параметр (список оценок) под названием xs. За параметром следует возвращаемый тип. 
В совокупности аргументы и возвращаемое значение также известны как сигнатура функции.
Наконец, далее идет тело функции, заключенное в фигурные скобки. 
В теле вызывается встроенная функция panic, которая вызывает ошибку выполнения (о ней я расскажу чуть позже в этой главе). 
Процесс написания функций может быть сложен, поэтому деление этого процесса на несколько частей вместо попытки реализовать всё за один большой шаг — хорошая идея.

Теперь давайте перенесём часть кода из функции main в функцию average:
package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	} 
	return total / float64(len(xs))
}
func main() {
	xs := []float64{98,93,77,82,83}
	
	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}
Обратите внимание, что мы заменили вызов fmt.Println на оператор return. 
Оператор возврата немедленно прервет выполнение функции и вернет значение, указанное после оператора, в функцию, которая вызвала текущую. 
Приведем main к следующему виду:

package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v:= range xs {
		total += v 
	}
	return total / float64(len(xs))
}	
func main() {
	xs := []float64{98,93,77,82,83}
	fmt.Println(average(xs))
}

Запуск этой программы должен дать точно такой же результат, что и раньше.

Несколько моментов, которые нужно иметь ввиду:

имена аргументов не обязательно должны совпадать с именами переменных при вызове функции. Например, можно сделать так:

func main() {
    someOtherName := []float64{98,93,77,82,83}
    fmt.Println(average(someOtherName))
}
и программа продолжит работать;

функции не имеют доступа к области видимости родительской функции, то есть это не сработает:

func f() {
    fmt.Println(x)
}
func main() {
    x := 5
    f()
}
Как минимум нужно сделать так:

func f(x int) {
    fmt.Println(x)
}
func main() {
    x := 5
    f(x)
}
или так:

var x int = 5
func f() {
    fmt.Println(x)
}
func main() {
    f()
}
функции выстраиваются в «стек вызовов». Предположим, у нас есть такая программа:

func main() {
    fmt.Println(f1())
}
func f1() int {
    return f2()
}
func f2() int {
    return 1
}
Её можно представить следующим образом:


Каждая вызываемая функция помещается в стек вызовов, каждый возврат из функции возвращает нас к предыдущей приостановленной подпрограмме;

можно также явно указать имя возвращаемого значения:

func f2() (r int) {
    r = 1
    return
}*/