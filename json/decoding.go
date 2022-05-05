// Функция Unmarshal декодирует данные формата json в Go структуру

// Декодирует json в структуру и срез структур go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Id         int
	Name       string
	Occupation string
}

func main() {
	var u1 user

	// Есть объект json, который преобразуется в байты
	data := []byte(`{
	 "Id": 1,
	 "Name" : "Valentin",
	 "Occupation": "programmer" 
 }`)

	// Декодируем структуру u1 из формата json в структуру через функцию Unmarshal
	err := json.Unmarshal(data, &u1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Struct is:", u1)
	fmt.Printf("%s is a %s. \n", u1.Name, u1.Occupation)

	// Декларируем срез структуры User
	var u2 []user

	// Данный массив json будет декодирован в срез
	data2 := []byte(`
[
	{"Id":2,"Name":"Vasiliy","Occupation":"driver"},
	{"Id":3,"Name":"Ivan","Occupation":"teacher"},
	{"Id":4,"Name":"Vladimir","Occupation":"president"}
]`)

	// Декодируем json-массив в срез через функцию Unmarshal
	err2 := json.Unmarshal(data2, &u2)
	if err2 != nil {
		log.Fatal(err2)
	}

	// Выводим декодированный json массив, пользователь за пользователем
	for i := range u2 {
		fmt.Println(u2[i])
	}

}
