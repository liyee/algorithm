package main

import "fmt"

//顺序表查找
//P⒤=1/n ASL=(n+1)/2
//P⒤=1/(2n) ASL=3(n+1)/4
func main() {
	table := []int{2, 3, 5, 1, 8, 4, 10, 100}
	key := 80

	exist := search(table, key)
	fmt.Println(exist)
}

func search(table []int, key int) bool {
	for _, num := range table {
		if num == key {
			return true
		}
	}

	return false
}
