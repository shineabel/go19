package main

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Detail struct{
		Home string `json:"home"`
		Phone string `json:"phone"`
	}  `json:"detail"`

}

var str = `{
"name":"shine",
"title":"manager",
"detail":{
  "home":"shanghai",
  "phone":"13888888888"
}

}`

func main() {

var c Contact

err := json.Unmarshal([]byte(str),&c)
if(err != nil){
	fmt.Println(err)
}
fmt.Println(c)
}
