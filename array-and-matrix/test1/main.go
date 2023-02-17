package main

import (
	"fmt"
	"sort"
)

//不重复打印排序数组中相加和为给定值的所有组合
func main() {
	arr := []int{2, 2, 3, 4, 5, 6}

	// UniqPair(arr, 10)
	// UniqTriad(arr, 10)
	combinationSum2(arr, 10)
}

//不重复打印排序数组中相加和为给定值的所有组合
var (
	res  [][]int
	path []int
)

func combinationSum2(arr []int, target int) {
	// res, path = make([][]int, 0), make([]int, 0, len(arr))
	sort.Ints(arr) // 排序，为剪枝做准备
	dfs(arr, 0, target)
	fmt.Println(res)
}

func dfs(arr []int, start int, target int) {
	if target == 0 { // target 不断减小，如果为0说明达到了目标值
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(arr); i++ {
		if arr[i] > target { // 剪枝，提前返回
			break
		}
		// i != start 限制了这不对深度遍历到达的此值去重
		if i != start && arr[i] == arr[i-1] { // 去重
			continue
		}
		path = append(path, arr[i])
		dfs(arr, i+1, target-arr[i])
		path = path[:len(path)-1]
	}
}

//3数组和
func UniqTriad(arr []int, k int) {
	if len(arr) < 3 {
		return
	}

	for i := 0; i < len(arr)-2; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			Rest(arr, i, i+1, len(arr)-1, k-arr[i])
		}
	}
}

func Rest(arr []int, f, l, r, k int) {
	for l < r {
		if arr[l]+arr[r] < k {
			l++
		} else if arr[l]+arr[r] > k {
			r--
		} else {
			// if l == f+1 || arr[l-1] != arr[l] {
			if arr[l-1] != arr[l] || l == f+1 {
				fmt.Println(arr[f], ",", arr[l], ",", arr[r])
				l++
				r--
			}
		}
	}
}

//2数组和
func UniqPair(arr []int, k int) {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[left]+arr[right] > k {
			right--
		} else if arr[left]+arr[right] < k {
			left++
		} else {
			if left == 0 || arr[left-1] != arr[left] {
				fmt.Println(arr[left], ", ", arr[right])
			}
			left++
			right--
		}
	}
}
