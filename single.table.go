package main

import "fmt"

type list struct {
	postion int
	next    *list
}

func (list *list) Append(node *list) error {

	//初始化判断
	if list == nil {

		list.next = nil // 声明一个空指针
		list.postion = node.postion
		return nil
	}

	//递归查找子节点
	origin := list
	for {

		if list.next == nil {
			list.next = node
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
