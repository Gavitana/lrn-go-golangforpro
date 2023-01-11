package main

import (
	"fmt"
	"github.com/Gavitana/lrn-go-golangforpro/ch01/utils"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	check := utils.CheckArgs(arguments)
	if !check {
		os.Exit(1)
	}
	var sum float64
	i := 1
	var numToDel int
	for ; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			n = sum + n
			sum = n
		} else {
			fmt.Println(err, ", skip this arg")
			numToDel++
		}

	}
	avg := sum / (float64(i-numToDel) - 1)
	fmt.Println("Average of float nums:", avg)
}
