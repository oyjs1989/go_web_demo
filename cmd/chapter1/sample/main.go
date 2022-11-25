package main

import (
	"log"
	"os"
	// _ "github.com/goinaction/code/chapter2/sample/mathchers"
	// "github.com/goinaction/code/chapter2/sample/"
)


func init(){
	log.SetOutput(os.Stdout)
}


func main() {
	search.Run("president")
}