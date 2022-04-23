package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type number struct {
	A int `json:"a"`
	B int `json:"b,omitempty"`
}

type add struct {
	C int `json:"c"`
}

type div struct {
	D float32 `json:"d"`
}


func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello\n")

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	err = nil
	var numb number
	err = json.Unmarshal(buf, &numb)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	fmt.Println("Наша структура", numb)
	fmt.Println(req.Method)
	fmt.Println(string(buf))

	// ___________________________

	//     json.Marshal()
	//     w.
	// // ____________________________

	// 	w.WriteHeader(200)
}

func bye(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Bye\n")

}

func addition(w http.ResponseWriter, req *http.Request) {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = nil
	var request number
	err = json.Unmarshal(buf, &request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	var response add
	response.C = request.A + request.B

	b, err := json.Marshal(&response)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	n, err := w.Write(b)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println("Количество считанных байт", n)
	fmt.Println(req.Method)
	fmt.Println(string(buf))

}

func subtraction(w http.ResponseWriter, req *http.Request) {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = nil
	var request number
	err = json.Unmarshal(buf, &request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	var response add
	response.C = request.A - request.B

	b, err := json.Marshal(&response)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	n, err := w.Write(b)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println("Количество считанных байт", n)
	fmt.Println(req.Method)
	fmt.Println(string(buf))
}
func multiplication(w http.ResponseWriter, req *http.Request) {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	err = nil
	var request number
	err = json.Unmarshal(buf, &request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	var response add
	response.C = request.A * request.B

	b, err := json.Marshal(&response)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	n, err := w.Write(b)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Println("Количество считанных байт", n)
	fmt.Println(req.Method)
	fmt.Println(string(buf))
}
func division(w http.ResponseWriter, req *http.Request) {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	err = nil
	var request number
	err = json.Unmarshal(buf, &request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	if request.B <= 0 {
		w.WriteHeader(400)
		return
	}

	var response div
	response.D = float32(request.A) / float32(request.B)

	b, err := json.Marshal(&response)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	n, err := w.Write(b)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Println("Количество записанных байтов", n)
	fmt.Println(req.Method)
	fmt.Println(string(buf))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}


func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/bye", bye)

	http.HandleFunc("/addition", addition)
	http.HandleFunc("/subtraction", subtraction)
	http.HandleFunc("/multiplication", multiplication)
	http.HandleFunc("/division", division)

	http.ListenAndServe(":8090", nil)
}
