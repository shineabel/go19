package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//func init() {
//	if len(os.Args) != 2 {
//		fmt.Println("url not provide")
//
//		os.Exit(1)
//	}
//}
func main() {
	//resp, err := http.Get(os.Args[1])
	//if(err != nil){
	//	fmt.Println("error:",err)
	//	return
	//}
	//
	//io.Copy(os.Stdout,resp.Body)
	//
	//if err = resp.Body.Close(); err != nil {
	//	fmt.Println("close error:",err)
	//}

	var b bytes.Buffer
	b.Write([]byte("hello"))

	fmt.Fprint(&b,", world")
	io.Copy(os.Stdout,&b)

}
