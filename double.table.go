package main

import (
	"fmt"
)

//声明链表结构体
type list struct {
	postion int
	last    *list //上一个节点指针
	next    *list //下一个节点指针

}

func (l *list) Append(node *list) error {

	//第一个元素初始化
	if l == nil {
		l.postion = node.postion
		l.last = nil
		l.next = nil
	}
	//递归查找
	for {

		if l.next == nil {

			node.last = l //记录上一个节点
			l.next = node
			return nil
		}

		l = l.next
	}

}
func main() {

	var l list
	for i := 1; i < 8; i++ {

		l.Append(&list{postion: i})

	}

	fmt.Printf("%v\n", l.next)

}
