package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if nil == root {
		return result
	}
	flag := false
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			temp[i] = queue[i].Val
			if nil != queue[i].Left {
				queue = append(queue, queue[i].Left)
			}
			if nil != queue[i].Right {
				queue = append(queue, queue[i].Right)
			}
		}
		if flag {
			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
		flag = !flag
		queue = queue[size:]
		result = append(result, temp)
	}
	return result
}
