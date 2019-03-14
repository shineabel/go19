package main

import (
	"log"
	"os"
	_ "github.com/go19/search"
	"github.com/go19/search"
	_ "github.com/go19/matcher"
)

func init()  {
	log.SetOutput(os.Stdout)
}
func main() {

	search.Run("Energy")
}
