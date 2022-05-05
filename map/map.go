// package main

// import "fmt"

// var userInfo = make(map[string]string)

// func main() {
// 	// userInfo := map[string]string {
// 	// 	"LastName": " Kovalyov",
// 	// 	"FirstName": " Valentin",
// 	// 	"MuddleName": " Maksimovic",
// 	// }

// 	last := userInfo[" LastName"]
// 	fmt.Println(last)

// 	userInfo["LastName"] = "Pupkin"
// 	fmt.Println(userInfo["LastName"])

// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	elements := map[string]map[string]string{
// 		"H": {
// 			"name":  "Hydrogen",
// 			"state": "gas",
// 		},
// 		"He": {
// 			"name":  "Helium",
// 			"state": "gas",
// 		},
// 		"Li": {
// 			"name":  "Lithium",
// 			"state": "solid",
// 		},
// 		"Be": {
// 			"name":  "Beryllium",
// 			"state": "solid",
// 		},
// 		"B": {
// 			"name":  "Boron",
// 			"state": "solid",
// 		},
// 		"C": {
// 			"name":  "Carbon",
// 			"state": "solid",
// 		},
// 		"N": {
// 			"name":  "Nitrogen",
// 			"state": "gas",
// 		},
// 		"O": {
// 			"name":  "Oxygen",
// 			"state": "gas",
// 		},
// 		"F": {
// 			"name":  "Fluorine",
// 			"state": "gas",
// 		},
// 		"Ne": {
// 			"name":  "Neon",
// 			"state": "gas",
// 		},
// 	}

// 	if el, ok := elements["F"]; ok {
// 		fmt.Println(el["name"], el["state"])
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type userInfo struct {
	LastName   string `json:"last,omitempty"`
	FirstName  string `json:"first,omitempty"`
	MiddleName string `json:"middle,omitempty"`
}

type filters struct {
	Id int `json:"id,omitempty"`
}

type createUserProfileResponse struct {
	Id string `json:"id,omitempty"`
}

var bd = make(map[string]userInfo)

func createUserProfile(w http.ResponseWriter, req *http.Request) {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	user := userInfo{}

	err = json.Unmarshal(buf, &user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	id := uuid.New()

	bd[id.String()] = user

	var response createUserProfileResponse

	response.Id = id.String()

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
	fmt.Println(n)

}

func listUser(w http.ResponseWriter, req *http.Request) {

	byteRes, err := json.Marshal(&bd)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	n, err := w.Write(byteRes)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	fmt.Println(n)

}

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

}

func main() {
	http.HandleFunc("/v1/user/name", createUserProfile)
	http.HandleFunc("/v1/user/name/list", listUser)
	http.HandleFunc("/v1/user/name/filter", filter)

	http.ListenAndServe(":8080", nil)
}
