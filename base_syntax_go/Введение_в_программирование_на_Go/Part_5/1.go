// Пример того, как обратиться к четвертому элементу массива и среза. т.к. нумерация массивов начинается с 0
package main

import "fmt"

func main() {
    var x [5]int
    x[3] = 100
    fmt.Println(x)
}