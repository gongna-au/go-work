package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Data struct {
	Name string
	Age  int
}
type Ret struct {
	Code  int
	Param string
	Msg   string
	Data  []Data
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	data := Data{Name: "why", Age: 18}

	ret := new(Ret)
	id := req.FormValue("id")
	//id := req.PostFormValue('id')

	ret.Code = 0
	ret.Param = id
	ret.Msg = "success"
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret_json, _ := json.Marshal(ret)

	io.WriteString(w, string(ret_json))
}
func HelloServer1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world1!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/hello1", HelloServer1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
