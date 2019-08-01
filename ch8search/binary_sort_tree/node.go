// Author: xufei
// Date: 2019-08-01

package binary_sort_tree

import (
	"fmt"
	"log"
)

func (n *BiTNode) replaceNode(parent, elem *BiTNode) error {
	if n == nil {
		return fmt.Errorf("replace node not allow nil")
	}
	if parent.left == n {
		parent.left = elem
	} else {
		parent.right = elem
	}
	return nil
}

// 中序打印
func Print(t *BiTNode) {
	if t == nil {
		return
	}

	log.Println(t.data) // 前序
	Print(t.left)
	//log.Println(t.data) // 中序
	Print(t.right)
	//log.Println(t.data) // 后序
}
