package main

import (
	"fmt"
	"sort"
)

//不重复打印排序数组中相加和为给定值的所有组合（含负数）
/*
	解一：整体加一个数，把所有负数变为正数来处理
	解二：如下
*/
func main() {
	arr := []int{-3, -2, -1, 0, 1, 2, 3}
	combinationSum2(arr, 4)
}

var (
	res  [][]int
	path []int
)

func combinationSum2(arr []int, target int) {
	sort.Ints(arr)
	dfs(arr, target, 0)
	fmt.Println(res)
}

func dfs(arr []int, target, start int) {
	if target == 0 && len(path) > 0 { //path长度为0时跳过
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		return
	}

	for i := start; i < len(arr); i++ {
		if arr[i] > target {
			break
		}

		if start != 0 && arr[i] == arr[i-1] {
			continue
		}

		path = append(path, arr[i])
		dfs(arr, target-arr[i], i+1)
		path = path[:len(path)-1]
	}
}
