package main

import "fmt"

//折半查找
//P⒤=1/n ASL=log(n+1)-1
func main() {
	table := []int{5, 13, 19, 21, 37, 56, 64, 75, 80, 88, 92}
	key := 19

	exist := search(table, key)

	fmt.Println(exist)
	fmt.Println(table[exist])
}

func search(table []int, key int) int {
	len := len(table)
	low, high := 0, len-1

	for low <= high {
		mid := (low + high) / 2
		if table[mid] == key {
			return mid
		} else if key < table[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0
}
