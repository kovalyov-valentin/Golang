/*Используя в качестве примера функцию makeEvenGenerator напишите makeOddGenerator, генерирующую нечётные числа.*/

package main 

import "fmt"

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return ret
    }
}
func main() {
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
	fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
	fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
	fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
	fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
    fmt.Println(nextOdd()) 
}
