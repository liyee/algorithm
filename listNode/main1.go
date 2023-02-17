package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
func main() {
	var head = new(ListNode)
	// head.data = 0
	var tail *ListNode
	tail = head //tail用于记录头结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = ListNode{Val: i}
		node.Next = tail //将新插入的node的next指向头结点
		tail = &node     //重新赋值头结点
	}
	Shownode(tail) //遍历结果
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	for tmp := dummyHead; tmp.Next != nil; {

	}

	return dummyHead.Next
}
