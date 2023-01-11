package utils

import (
	"errors"
	"fmt"
	"strconv"
)

func CheckArgs(args []string) bool {
	// check that int args exists
	var err error = errors.New("An error")
	k := 1
	for err != nil {
		if k >= len(args) {
			fmt.Println("None of the arguments is a int!")
			return false
		}
		_, err = strconv.ParseInt(args[k], 0, 64)
		k++
	}
	return true

}
