package main

import (
	"fmt"
)

type mapData map[string]int

func main() {
	arr := []string{"abc", "ab", "a"}
	res := f1(arr)
	fmt.Println((res))
}

func f1(arr []string) mapData {
	nums := make(map[string]int, 1)

	for _, val := range arr {
		for _, letter := range val {
			nums[string(letter)]++
		}
	}

	return mapData(nums)
}
