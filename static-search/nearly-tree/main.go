package main

import (
	"fmt"
	"math"
)

//静态树表的查找-次优查找树
func main() {
	bt := BTreeNode{}
	list := []Element{
		{data: "A", weight: 1},
		{data: "B", weight: 1},
		{data: "C", weight: 2},
		{data: "D", weight: 5},
		{data: "E", weight: 3},
		{data: "F", weight: 4},
		{data: "G", weight: 4},
		{data: "H", weight: 3},
		{data: "I", weight: 5},
	}
	sw := []float64{0}
	for i := 0; i < len(list); i++ {
		sw = append(sw, list[i].weight+sw[i])
	}

	low := 1
	high := len(list)
	SecondOptimal(&bt, list, sw, low, high)

	fmt.Println(bt)
}

type BTreeNode struct {
	data           string
	lchild, rchild *BTreeNode
}

type Element struct {
	data   string
	weight float64
}

func SecondOptimal(BT *BTreeNode, list []Element, sw []float64, low int, high int) {
	i := low
	min := math.Abs(sw[high] - sw[low])
	dw := sw[high] + sw[low-1]

	for j := low + 1; j <= high; j++ {
		if math.Abs(dw-sw[j]-sw[j-1]) < min {
			i = j
			min = math.Abs(dw - sw[j] - sw[j-1])
		}
	}

	BT.data = list[i-1].data
	if i == low {
		BT.lchild = nil
	} else {
		//构造左子树
		if BT.lchild == nil {
			BT.lchild = &BTreeNode{}
		}
		SecondOptimal(BT.lchild, list, sw, low, i-1)
	}
	if i == high {
		BT.rchild = nil
	} else {
		if BT.rchild == nil {
			BT.rchild = &BTreeNode{}
		}
		SecondOptimal(BT.rchild, list, sw, i+1, high)
	}
}
