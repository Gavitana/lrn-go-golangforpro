package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"time"
)

func fibonacci(a int) int {
	e := 1
	var c, b int
	for i := 0; i < a; i++ {
		b = c + e
		c = e
		e = b
	}
	return b
}

func main() {
	fmt.Println("Now we check speed of fibonacci sequence function in C and Golang")
	fmt.Println("We take average of 10000000 iteration of function")
	fmt.Println("Starts with Golang")
	allRes := time.Duration(0)
	for i := 0; i < 10000000; i++ {
		start := time.Now()
		_ = fibonacci(1000)
		elapsed := time.Since(start)
		allRes += elapsed

	}
	goAvg := allRes / 10000000
	fmt.Println()
	fmt.Printf("Golang fibonacci function time = %s", goAvg.String())
	fmt.Println()
	fmt.Println("Now C")
	allRes2 := time.Duration(0)
	for i := 0; i < 10000000; i++ {
		start := time.Now()
		_ = C.fibonacci(1000)
		elapsed := time.Since(start)
		allRes2 += elapsed
	}
	goAvg1 := allRes / 10000000
	fmt.Printf("C fibonacci function time = %s", goAvg1.String())
}
