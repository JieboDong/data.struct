package main

import "fmt"

type list struct {
	postion int
	next    *list
}

func (list *list) Append(node *list) error {

	//初始化判断
	if list == nil {

		list.next = nil             // 同时是单链表的尾部
		list.postion = node.postion // 单链表有了第一个元素
		return nil
	}

	//递归查找子节点
	for {
		if list.next == nil {
			list.next = node // 同时是单链表的尾部
			return nil
		}

		list = list.next

	}

}

func main() {

	//初始化
	var l list
	for i := 1; i < 10; i++ {

		n := list{postion: i}
		l.Append(&n)

	}
	fmt.Printf("%v\n", l)

}
