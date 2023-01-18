// Функция Marshal кодирует данные в формат JSON
// В данном примере мы трансформируем структуру и срез структуры в формат JSON

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Объявляем структуру User
type User struct {
	Id         int
	Name       string
	Occupation string
}

func main() {

	// Создаем экземпляр стуктуры
	u1 := User{1, "Valentin", "programmer"}

	// Кодируем структуру u1 в JSON с помощью функции Marshal
	jsonData, err := json.Marshal(u1)
	if err != nil {
		log.Fatal(err)
	}

	// Так как jsonData является байтовым массивом, мы конвертируем его в строку с помощью строковой функции
	fmt.Println(string(jsonData))

	// У нас есть срез users
	users := []User{
		{Id: 2, Name: "Vasiliy", Occupation: "driver"},
		{Id: 3, Name: "Ivan", Occupation: "teacher"},
		{Id: 4, Name: "Anton", Occupation: "manager"},
	}

	// Кодируем срез users через функцию Marshal
	jsonData2, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	// Выводим закодированный срез
	fmt.Println(string(jsonData2))
}
