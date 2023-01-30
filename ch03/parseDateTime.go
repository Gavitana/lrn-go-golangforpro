package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myTime string
	var myDate string
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	myTime = os.Args[2]
	myDate = os.Args[1]
	t, err := time.Parse("15:04", myTime)
	d, err := time.Parse("02 Jan 2006", myDate)
	if err == nil {
		fmt.Println("Day:", d.Day())
		fmt.Println("Month:", d.Month())
		fmt.Println("Year:", d.Year())
		fmt.Println("Hour:", t.Hour())
		fmt.Println("Minutes:", t.Minute())
	} else {
		fmt.Println(err)
	}
}
