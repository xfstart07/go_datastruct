// Author: xufei
// Date: 2019-08-01

package avl

import print2 "datastruct/ch8search/print"

func (n *BitNode) getHeight() int {
	if n == nil {
		return 0
	}

	return n.height
}

func (n *BitNode) MaxHeight() int {
	if n.left.getHeight() > n.right.getHeight() {
		return n.left.getHeight()
	}

	return n.right.getHeight()

}

func (n *BitNode) MinHeight() int {
	if n.left.getHeight() < n.right.getHeight() {
		return n.left.getHeight()
	}
	return n.right.getHeight()
}

// 左旋
func (n *BitNode) LeftRotate() *BitNode {
	r := n.right
	n.right = r.left
	r.left = n

	n.updateHeight()
	r.updateHeight()

	return r
}

// 右旋
func (n *BitNode) RightRotate() *BitNode {
	l := n.left
	n.left = l.right
	l.right = n

	n.updateHeight()
	l.updateHeight()

	return l
}

// 先左旋后右旋
func (n *BitNode) LeftThenRightRotate() *BitNode {
	n.left = n.left.LeftRotate()
	return n.RightRotate()
}

// 先右旋后左旋
func (n *BitNode) RightThenLeftRotate() *BitNode {
	n.right = n.right.RightRotate()
	return n.LeftRotate()
}

func (n *BitNode) updateHeight() {
	n.height = n.MaxHeight() + 1
}

func (n *BitNode) Data() interface{} {
	return n.data
}

func (n *BitNode) LeftNode() print2.Printer {
	return n.left
}

func (n *BitNode) RightNode() print2.Printer {
	return n.right
}
