package leetcode

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if nil == root {
		return result
	}
	queue1 := list.New()
	queue1.PushBack(root)
	for nil != queue1.Front() {
		size := queue1.Len()
		temp := make([]int, 0)
		for size > 0 {
			node := queue1.Remove(queue1.Front()).(*TreeNode)
			if nil != node.Left {
				queue1.PushBack(node.Left)
			}
			if nil != node.Right {
				queue1.PushBack(node.Right)
			}
			temp = append(temp, node.Val)
			size--
		}
		result = append(result, temp)
	}
	return result
}
