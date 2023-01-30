package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {

	var myDate string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}

	myDate = os.Args[1]
	d, err := time.Parse("02-Jan-2006", myDate)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Day:", d.Day())
		fmt.Println("Month:", d.Month())
		fmt.Println("Year:", d.Year())
	} else {
		fmt.Println(err)
	}
}
