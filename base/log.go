package main

import (
	"log"
	"os"
	"io/ioutil"
	"io"
)

//func init() {
//	log.SetPrefix("TRACE:")
//	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
//
//}

func init() {
	file, err := os.OpenFile("error.txt",os.O_CREATE | os.O_RDONLY | os.O_APPEND,0666)
	if(err != nil){
		log.Fatalln("failed to open file:",err)
	}
	Trace = log.New(ioutil.Discard,"Trace:",log.Ldate | log.Lmicroseconds | log.Llongfile)
	Info = log.New(os.Stdout,"Info:",log.Ldate | log.Lmicroseconds | log.Llongfile)
	Warn = log.New(os.Stdout,"Warn:",log.Ldate | log.Lmicroseconds | log.Llongfile)
	Error = log.New(io.MultiWriter(file,os.Stderr),"Error:",log.Ldate | log.Lmicroseconds | log.Llongfile)
}

var (


	Trace *log.Logger
	Info *log.Logger
	Warn *log.Logger
	Error *log.Logger
)
func main() {
	//log.Println("message")
	//log.Fatalln("Fatal message")
	//log.Panicln("panic message")

	Trace.Println("I have something to say")
	Info.Println("special message")
	Warn.Println("important information you need to know")
	Error.Println("something failed")
}
