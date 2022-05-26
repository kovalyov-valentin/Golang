// package main

// import (	
// 	"fmt"		
// )

// type userInfo struct {
// 	LastName   string `json:"last,omitempty"`
// 	FirstName  string `json:"first,omitempty"`
// 	MiddleName string `json:"middle,omitempty"`
// }

// // type filters struct {
// // 	Id int `json:"id,omitempty"`
// // }

// var bd = make(map[int]userInfo)

// func main() {
//     n := len(bd)
//     fmt.Println(n)
// }
package main 

import (
	"fmt"
)
func decode(encoded []int, first int) []int {
    arr := make([]int, len(encoded) + 1)
    arr[0] = first 
    for i, v := range encoded {
        arr[i + 1] = arr[i] ^ v
    }
    return arr
}
func main() {
	dec := 
}