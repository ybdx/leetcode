package leetcode

func generateTrees(n int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if 0 == n {
		return result
	}
	return gen_trees(1, n)
}

func gen_trees(start, end int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if start > end {
		result = append(result, nil)
		return result
	}
	for i:= start; i<=end; i++ {
		leftTrees := gen_trees(start, i-1)
		rightTrees := gen_trees(i+1, end)
		for _, l := range leftTrees {
			for _, r:= range rightTrees {
				curNode := &TreeNode{Val:i}
				curNode.Left = l
				curNode.Right = r
				result = append(result, curNode)
			}
		}
	}
	return result
}
