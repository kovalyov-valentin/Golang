package main

import (	
	"fmt"		
)

type userInfo struct {
	LastName   string `json:"last,omitempty"`
	FirstName  string `json:"first,omitempty"`
	MiddleName string `json:"middle,omitempty"`
}

// type filters struct {
// 	Id int `json:"id,omitempty"`
// }

var bd = make(map[int]userInfo)

func main() {
    n := len(bd)
    fmt.Println(n)
}