package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
)

//Стурктура запроса
type request struct {

	// Переменная запроса вида
	// "1+1"
	Req string `json:"req"`
}

//Структура ответа
type response struct {
	
	//Переменная ответа вида
	//"1+1=2"
	Res string `json:"res"`
}

// Функция обработчика для пути "/v1/calculator"
func calculator(w http.ResponseWriter, req *http.Request) {
	
	// Читает тело запроса
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Создаем переменную типа request
	var requ request 

	// Преводим байты в нужную нам структуру (анмаршалим)
	err = json.Unmarshal(buf, &requ)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	
	//Создаем переменную типа rune 
	var operator rune
	
	// Обрабатываем возможную ошибку запроса
	if len(requ.Req) == 0 {
		w.WriteHeader(404)
		return 
	} 
	
	// Находим оператор (+,-,*)
	for i := 0; i < len(requ.Req); i++ {
		switch rune(requ.Req[i]) {
		case '+':
			operator = '+'
		case '-':
			operator = '-'
		case '*':
			operator = '*'
		case '/':
			operator = '/'
		default:
			w.WriteHeader(404)
			return
		}
	}

	// Делим запрос и получаем массив строк, который равен нужным нам числам
	str := strings.Split(requ.Req, string(operator))
		fmt.Println(str[0], str[1])

	// Из того, что мы спарсили получаем число а из string
	a, err := strconv.Atoi(str[0])
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// Из того, что мы спарсили получаем число b из string
	b, err := strconv.Atoi(str[1])
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// Создаем переменную суммы двух чисел, которыемы получили sum = a + b
	var sum int
	
	//Operator определяет нужный нам математический оператор
	switch operator {
	case '+':
		sum = a + b
	case '-':
		sum = a - b 
	case '*':
		sum = a * b
	case '/':
		if b == 0 {
			w.WriteHeader(400)
			return
		}
		sum = a / b
	}
	

	// Создаем переменную результата наших строк
	var resStr string 
	
	// Создаем переменную для перевода нашей суммы из инта в стринг 
	var sumStr string
	
	//Функция переводит сумму из инт в стринг 
	sumStr = strconv.Itoa(sum)
	
	// Конкатенация строк(слияние)
	resStr = requ.Req + "=" + sumStr

	// Определяем переменную ответа для нашего клиента типа response
	var res response
	
	// Присваием нашему полю в структуре результат слияния строк 
	res.Res = resStr

	// Переводим структуру в байты(маршалим)
	byteRes, err := json.Marshal(&res)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	//Записываем ответы в http ответы(в body)
	n, err := w.Write(byteRes)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println("Количество записанных байт", n)
}

func main() {
 http.HandleFunc("/v1/calculator", calculator)
 http.ListenAndServe(":8080", nil)
}