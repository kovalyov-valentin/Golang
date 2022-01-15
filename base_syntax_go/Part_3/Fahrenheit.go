package main

import "fmt"

func main() {
    fmt.Print("Enter the temperature in degrees Fahrenheit: ")
    var F float64
    fmt.Scanf("%f", &F)

    C := (F - 32) * 5/9

    fmt.Println(C)    
}