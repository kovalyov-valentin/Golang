package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
)

type request struct {
	Req string `json:"req"`
}
type response struct {
	Res string `json:"res"`
}

func calculator(w http.ResponseWriter, req *http.Request) {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = nil
	var requ request 
	err = json.Unmarshal(buf, &requ)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	
	fmt.Println(requ)

	var operator rune
	if len(requ.Req) == 0{
		return 
	} 

	for i := 0; i < len(requ.Req); i++ {
		switch rune(requ.Req[i]) {
		case '+':
			operator = '+'
		case '-':
			operator = '-'
		case '*':
			operator = '*'
		}
	}

	str := strings.Split(requ.Req, string(operator))
		fmt.Println(str[0], str[1])

	a, err := strconv.Atoi(str[0])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	b, err := strconv.Atoi(str[1])
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println(a, b)
	var sum int
	switch operator {
	case '+':
		sum = a + b
	case '-':
		sum = a - b 
	case '*':
		sum = a * b
	}
	
	fmt.Println(requ.Req, "=", sum)
	var resStr string 
	
	var sumStr string
	
	sumStr = strconv.Itoa(sum)
	
	resStr = requ.Req + "=" + sumStr

	fmt.Println(resStr)

	var res response
	res.Res = resStr
	fmt.Println("qjfgj", res.Res)
	byteRes, err := json.Marshal(&res)
	if err != nil {
		w.WriteHeader(500)
		return
	}
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