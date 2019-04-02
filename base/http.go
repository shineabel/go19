package main

import (
	"net/http"
	"log"
	"fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Printf("%s %s\n",r.URL.Path,r.Method)
	w.Write([]byte("hello"))


}
func main() {

	http.HandleFunc("/",helloHandler)

	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal("server error")
	}
}
