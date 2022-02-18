// Go имеет встроенный тип для сообщение об ошибках(тип error).
// Мы можем создавать свои собственные типы сообщений об ошибках, используя для этого функцию New из пакета errors.

package main

import (
	// "fmt"
	"errors"
)
func main() {
	err := errors.New("error message")
	// fmt.Println(err)
}