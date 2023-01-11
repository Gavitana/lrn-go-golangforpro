package main

import (
	"fmt"
	"github.com/Gavitana/lrn-go-golangforpro/ch01/utils"
	"os"
	"strconv"
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
	// sum args
	var sum int64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseInt(arguments[i], 0, 64)
		if err == nil {
			n = sum + n
			sum = n
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("Sum of int nums:", sum)
}
