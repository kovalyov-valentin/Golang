package main

import "fmt"

func main() {
    fmt.Print("Enter to convert to meters: ")
    var ft float64
    fmt.Scanf("%f", &ft)

    m := ft * 0.3048

    fmt.Println(m)    
}