package leetcode

import (
	"container/list"
)

func flatten(root *TreeNode)  {
	if nil == root {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		t := queue.Back().Value.(*TreeNode)
		queue.Remove(queue.Back())
		if nil != t.Left {
			r := t.Left
			for nil != r.Right {
				r = r.Right
			}
			r.Right = t.Right
			t.Right = t.Left
			t.Left = nil
		}
		if nil != t.Right {
			queue.PushBack(t.Right)
		}
	}
}
