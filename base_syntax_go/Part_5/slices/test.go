//Вот пример среза:
//var x []float64

//Срез создается встроенной функцией make:
//x := make([]float64, 5)
//Этот код создаст срез, который связан с массивом типа float64, длиной 5. 
//Срезы всегда связаны с каким-нибудь массивом. Они не могут стать больше чем массив, а вот меньше — пожалуйста. Функция make принимает и третий параметр:
//x := make([]float64, 5, 10)
//10 — это длина массива, на который указывает срез:

//Другой способ создать срез — использовать выражение [low : high]:
//arr := [5]float64{1,2,3,4,5}
//x := arr[0:5]
//low - это позиция, с которой будет начинаться срез, а high - это позиция, где он закончится. 
//Например: arr[0:5] вернет [1,2,3,4,5], arr[1:4] вернет [2,3,4].
//Для удобства мы также можем опустить low, high или и то, и другое. arr[0:] это то же самое что arr[0:len(arr)], 
//arr[:5] то же самое что arr[0:5] и arr[:] то же самое что arr[0:len(arr)].



package main

import "fmt"

func main() {
	x := make([]float64, 5)
}