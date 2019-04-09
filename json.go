package main

import (
	"fmt"
	"encoding/json"
)

func main() {

	b := []byte(
		`{
"Title":"Go",
"Authors":["shine","cat"],
"FlPublished":true,
"Price":35.85,
"SaleCount":1000
		}`)

	fmt.Println(string(b))

	var r interface{}
	err := json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println("json ummarshal error", err)
	}

	book , ok := r.(map[string]interface{})
	if !ok {
		fmt.Println("json unmarshal type error")
	}

	for k , v := range book {
		switch  t := v.(type){
			case string:
				fmt.Printf("%s is string ,value:%s\n", k, t)
			case int:
				fmt.Printf("%s is int ,value:%s\n", k, t)
			case bool:
				fmt.Printf("%s is bool ,value:\n", k, t)
		case []interface{}:
			fmt.Printf("%s is an array\n", k)
			for index, value := range  t {
				fmt.Println(index, value)
			}
		default:
			fmt.Println("unknow type")
		}
	}
}
