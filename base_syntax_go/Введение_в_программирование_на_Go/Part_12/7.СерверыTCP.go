// На Го просто создавать сетевые серверы.
// Пример того, как создать TCP сервер:

package main

import (
	"encoding/gob"
	"fmt"
	"net"
)
func server() {
	//слушать порт
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		//принятие соединения
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//обработка соединения
		go handleServerConnection(c)
	}
}
func handleServerConnection(c net.Conn) {
	//получение сообщения
	var msg string 
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}
func client() {
	//соединиться с сервером
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//послать сообщение
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}
func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

// Этот пример использует пакет encoding/gob, который позволяет легко кодировать выходные данные, чтобы другие программы на Го могли их прочитать.
// Дополнительные способы кодирования доступны в пакете encoding