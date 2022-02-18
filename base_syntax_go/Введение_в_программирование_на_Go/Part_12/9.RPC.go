// Пакеты net/rpc (remote procedure call - удаленный вызов процедур) и net/rpc/jsonrpc обеспечивают простоту вызова методов по стети ( а не только из программы, в которой они используются)

package main

import (
	"fmt"
	"net"
	"net/rpc"
)
type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}
func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
func client() {
	c, err := rpc.Dial("tcp","127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64 
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}
func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

Эта программа похожа на пример использования TCP-сервера, за исключением того, что теперьмы создали объект, который содержит методы, доступные для вызова, а затем вызвали Negate из функции клиента. 
Для получения дополнительной информации стоит посмотреть документацию по net/rcp.