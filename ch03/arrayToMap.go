package main

import "fmt"

func arrayToHash(arr []int) map[int]int {
	iMap := make(map[int]int)
	for id, num := range arr {
		iMap[id] = num
	}
	return iMap
}

func main() {
	arr := []int{1, 3, 4, 5}
	for _, num := range arr {
		fmt.Println(num)
	}
	iMap := arrayToHash(arr)
	for key, value := range iMap {
		fmt.Println(key, value)
	}
}
