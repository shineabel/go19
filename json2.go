package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var r map[string]interface{}
		if err := dec.Decode(&r); err != nil {
			log.Println(err)
			return
		}
		for k := range r {
			if k != "Title" {
				delete(r, k)
			}
		}
		if err2 := enc.Encode(&r); err2 != nil {
			fmt.Println(err2)
		}
	}
}
