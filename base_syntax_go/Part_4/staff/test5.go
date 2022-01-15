package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
    var i float64
    fmt.Scanf("%f", &i) 
	
	i := 10

	if i > 10 {
    fmt.Println("Big")
	} else {
    fmt.Println("Small")
	}
}