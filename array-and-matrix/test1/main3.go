package main

import (
	"fmt"
)

//目标位k的组合数量（数组有正负数）

// func main() {
// 	arr := []int{2, -2, 3, 0, -7, 4}

// 	sort.Ints(arr) //[-7 -2 0 2 3 4]
// 	// fmt.Println(arr)

// 	num := Solution(arr)

// 	fmt.Println(num)
// }

func Solution(arr []int) int {
	num, sum := 0, 0
	l := len(arr)
	zero := false

	for i := 0; i < l-1; i++ {
		tmp := []int{arr[i]}
		sum = arr[i]
		if arr[i] == 0 {
			fmt.Println(tmp)
			num += 1
			break
		}

		for j := i + 1; j < l; j++ {
			if arr[j] == 0 {
				zero = true
			}
			tmp = append(tmp, arr[j])
			sum += arr[j]
			if sum == 0 {
				fmt.Println(tmp)
				if zero == true && len(tmp) > 1 {
					num += 1
				}
				num += 1
				break
			}
		}
	}

	return num
}
