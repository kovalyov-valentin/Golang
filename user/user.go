package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Структура запроса
type name struct {
	LastName   string `json:"last"`
	FirstName  string `json:"first"`
	MiddleName string `json:"middle"`
	Age        int    `json:"age"`
}

type filters struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Объявляем глобальную переменную, состоящую из срезов типа name
var bd []name

// Функция обработчика для пути /v1/user/name
func user(w http.ResponseWriter, req *http.Request) {

	// Читает тело запроса
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Создаем переменную типа name
	var userProfile name

	// Переводим байты в нужную нам структуру
	err = json.Unmarshal(buf, &userProfile)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	// fmt.Println(userProfile)

	bd = append(bd, userProfile)
	fmt.Println(bd)

}
func listUser(w http.ResponseWriter, req *http.Request) {

	// Переводим структуру в байты(маршалим)
	byteRes, err := json.Marshal(&bd)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// Записываем ответы в http ответы (в body)
	n, err := w.Write(byteRes)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println(n)
}

// В этой функции фильтруем базу данных
func filter(w http.ResponseWriter, req *http.Request) {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var fil filters

	err = json.Unmarshal(buf, &fil)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	for i := 0; i < len(bd); i++ {
		if bd[i].LastName == fil.Name {
			fmt.Println(bd[i])
		}
	}
}

func main() {
	http.HandleFunc("/v1/user/name", user)
	http.HandleFunc("/v1/user/name/list", listUser)
	http.HandleFunc("/v1/user/name/filter", filter)

	http.ListenAndServe(":8080", nil)
}
