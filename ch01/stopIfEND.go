package main

import (
	"fmt"
	"github.com/Gavitana/lrn-go-golangforpro/ch01/utils"
	"os"
)

func main() {
	// check that you gave an args
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more int.")
		os.Exit(1)
	}
	// get args to new var
	arguments := os.Args
	// check that int var exists
	check := utils.CheckArgs(arguments)
	if !check {
		os.Exit(1)
	}
	for i := 1; i < len(arguments); i++ {
		if arguments[i] == "END" {
			fmt.Println("END")
			return
		} else {
			fmt.Println(arguments[i])
		}
	}

}
