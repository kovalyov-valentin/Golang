// Вывод ошибок
// Представлен метод передачи данных в стандартный поток ошибок, который помогает различить реальные значения и вывод сообщений об ошибках.

package main 

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args 
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, "This is Standart output\n")
	io.WriteString(os.Stderr, myString) // Функция io.WriteString() вызывается два раза для записи в стандартный поток ошибок (os.Stderr) и еще один раз - для записи в стандартный поток вывода (os.Stdout)
	io.WriteString(os.Stderr, "\n")
}