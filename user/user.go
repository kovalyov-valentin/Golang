package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Структура запроса
type userProfile struct {
	LastName   string `json:"last"`
	FirstName  string `json:"first"`
	MiddleName string `json:"middle"`
	Age        int    `json:"age"`
}

type filters struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Объявляем глобальную переменную, состоящую из срезов типа userProfile
var bd []userProfile

// Функция обработчика для пути /v1/user/name
func user(w http.ResponseWriter, req *http.Request) {

	// Читает тело запроса
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Создаем переменную типа userProfile
	var user userProfile

	// Переводим байты в нужную нам структуру
	err = json.Unmarshal(buf, &user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	// fmt.Println(user)

	bd = append(bd, user)
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
			
			b, err := json.Marshal(bd[i])
			if err != nil {
				w.WriteHeader(500)
				return
			}

			n, err := w.Write(b)
			if err != nil {
				w.WriteHeader(500)
				return
			}

			fmt.Println(n)
			return
		}

	}

	for i := 0; i < len(bd); i++ {
		if i == fil.Id {

			b, err := json.Marshal(bd[i])
			if err != nil {
				w.WriteHeader(500)
				return
			}
			n, err := w.Write(b)
			if err != nil {
				w.WriteHeader(500)
				return
			}

			fmt.Println(n)
			return
		}

	}
}

func main() {
	http.HandleFunc("/v1/user/name", user)
	http.HandleFunc("/v1/user/name/list", listUser)
	http.HandleFunc("/v1/user/name/filter", filter)

	http.ListenAndServe(":8080", nil)
}
