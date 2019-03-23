package main

import (
	"os"
	"fmt"
	"io"
	"net/http"
	"log"
)

//func init() {
//	if len(os.Args) != 2 {
//		fmt.Println("url not provide")
//
//		os.Exit(1)
//	}
//}
func main() {
	resp, err := http.Get(os.Args[1])
	if(err != nil){
		fmt.Println("error:",err)
		return
	}
	f, err2 := os.Create(os.Args[2])
	if(err2 != nil){
		fmt.Println(err2)
	}

	dest := io.MultiWriter(os.Stdout,f)
	io.Copy(dest,resp.Body)
	if err3 := resp.Body.Close();err3 != nil {
		log.Println(err3)
	}



	//var b bytes.Buffer
	//b.Write([]byte("hello"))
	//
	//fmt.Fprint(&b,", world")
	//io.Copy(os.Stdout,&b)

}
