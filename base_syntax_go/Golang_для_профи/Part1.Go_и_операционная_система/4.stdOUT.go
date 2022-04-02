package main 

import (
	"io" // используется вместо пакета fmt
	"os" // Этот пакет применяется для чтения программой аргументов командной строки и для доступа к os.Stdout
)
func main() {
	myString := "" // Данная переменная содержит текст, который будет выведен на экран. 
	arguments := os.Args // Предоставляет доступ к аргументам командной строки Go
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString) /* В данном случае io.WriteString() работает также как fmt.Print() однако принимает только два параметра
	Первый из них - это файл, в который будут записываться результаты (здесь os.Stdout), второй строкая переменная */
	
	io.WriteString(os.Stdout, "\n")
}