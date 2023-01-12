package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var LOGFILE = "/tmp/mGo.log"
var LOGFILE2 = "/tmp/mGo2.log"

func main() {
	f1, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f2, err2 := os.OpenFile(LOGFILE2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer f1.Close()
	defer f2.Close()
	w := io.MultiWriter(os.Stdout, f1, f2)
	logger := log.New(w, "logger", log.LstdFlags)
	logger.Println("Same text in different files")
}
