package main

import (
	"fmt"
	"sort"
)

//分发饼干 https://leetcode.cn/problems/assign-cookies/
func main() {
	num := findContentChildren([]int{1, 2, 3}, []int{1, 2, 3})
	fmt.Println(num)
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	m, num := 0, 0
	for i := 0; i < len(g); i++ {
		for ; m < len(s); m++ {
			if g[i] <= s[m] {
				num++
				m++
				break
			}
		}
	}
	return num
}
