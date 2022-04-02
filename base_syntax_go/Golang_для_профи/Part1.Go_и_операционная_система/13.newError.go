// Тип данных error
// Пример программы для создания переменныз типа error
// Для создания переменной типа error необходимо вызвать функцию New() из стандартного пакета errors.
package main 

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!") // Функция errors.New() создает переменную типа error. Также данная функция принимает в качестве параметра значние типа string.
		return err
	} else {
		return nil // Если функция должна вернуть переменную типа error, но ошибка не произошла, то возвращается nil.
	}
}

func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normaly!")
	} else {
		fmt.Println(err)
	}
	
	if err.Error() == "Error in returnError() function!" { // Функция err.Errors() позволяет преобразовать переменную типа error в тип string.
		// Эта функция позволяет сравнивать переменную типа error с переменной типа string.

		fmt.Println("!!")
	}
}
// Как видно большая часть кода функции main() уходит на то, чтобы проверить, равна ли перменная типа error со значением nil, после чего можно действовать соответственно.

// Также следует знать что на UNIX машинах рекомендуется передвать сообщения об ошибках в службу журналировани, особенно если программа является сервером или важным приложением.