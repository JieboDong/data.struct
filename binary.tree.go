package main

import (
	"fmt"
)

//二叉树 结构声明
type tree struct {
	id    int
	left  *tree
	right *tree
}

//添加左节点
func (t *tree) Left(id int, node *tree) {

	if t == nil {

		return //如果没有该节点 就退出

	} else if t.id == id {

		if t.left == nil {
			t.left = node
			return
		}

	}

	t.left.Left(id, node)
	t.right.Right(id, node)
}

//添加右节点
func (t *tree) Right(id int, node *tree) {

	if t == nil {

		return //如果没有该节点 就退出

	} else if t.id == id {

		if t.right == nil {

			t.right = node
			return
		}
	}

	t.left.Left(id, node)
	t.right.Right(id, node)
}

func (l *tree) output() {

	var ll interface{}
	var rr interface{}
	if l.left == nil {

		ll = "不存在"
	} else {

		ll = l.left.id
	}

	if l.right == nil {

		rr = "不存在"
	} else {

		rr = l.right.id
	}

	fmt.Printf("父节点id：%v,右节点%v,左节点%v\n", l.id, rr, ll)

	if l.right != nil {

		l.right.output()

	}
	if l.left != nil {

		l.left.output()

	}

}

//初始化根结点
func (t *tree) Init() {

	t.id = 1
	t.left = nil
	t.right = nil

}

func main() {

	//初始化根节点
	var t tree
	t.Right(0, &tree{id: 9})
	t.Left(0, &tree{id: 10})
	t.Left(10, &tree{id: 11})

	fmt.Printf("%v", t.left.left.id)
	t.output()
}
