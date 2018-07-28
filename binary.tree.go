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
func (t *tree) Left(id int, node *tree, guide string) {

	if t == nil {

		return //如果没有该节点 就退出

	} else if t.id == id {

		if t.left == nil && guide == "left" {
			t.left = node
			return
		} else if t.right == nil && guide == "right" {
			t.right = node
			return

		}

	}

	t.left.Left(id, node, guide)
	t.right.Right(id, node, guide)
}

//添加右节点
func (t *tree) Right(id int, node *tree, guide string) {

	if t == nil {

		return //如果没有该节点 就退出

	}
	if t.id == id {

		if t.left == nil && guide == "left" {
			t.left = node
			return
		} else if t.right == nil && guide == "right" {
			t.right = node
			return

		}
	}
	fmt.Printf("%v\n", t)
	t.left.Left(id, node, guide)
	t.right.Right(id, node, guide)
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

//反转二叉树
func (t *tree) reversal() {

	//左节点不存在
	if t.left == nil && t.right != nil {

		l := t.right
		t.right = t.left
		t.left = l
		t.right.reversal()
		return
	}
	//右节点不存在
	if t.left != nil && t.right == nil {
		r := t.left
		t.right = r
		t.left = nil
		t.left.reversal()
		return
	}

	if t.left != nil && t.right != nil {
		r := t.left
		t.left = t.right
		t.right = r
		t.left.reversal()
		t.right.reversal()
		return
	}
	return
}

func main() {

	//初始化根节点
	var t tree
	t.Right(0, &tree{id: 9}, "right")
	t.Left(0, &tree{id: 10}, "left")

	t.Right(10, &tree{id: 11}, "right")
	t.Left(10, &tree{id: 12}, "left")
	t.Right(11, &tree{id: 23}, "right")
	t.Left(11, &tree{id: 24}, "left")
	fmt.Printf("%v\n", t)
	t.output()
	// t.reversal()
	fmt.Println("反转二叉树后排列\n")
	t.output()
}
