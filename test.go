package main
 
import "fmt"
 
// kelvinToCelsius конвертирует °K в °C
func kelvinToCelsius(k float64) float64 { // Объявляет функцию, что принимает параметр и возвращает результат
    k -= 273.15
    return k
}
 
func main() {
    kelvin := 294.0
    celsius := kelvinToCelsius(kelvin) // Вызывает функцию передачи kelvin как первого аргумента
    fmt.Println(kelvin, "° K is ", celsius, "° C") // Выводит: 294° K is 20.850000000000023° C
 }