package leetcode

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	return visit(root, &sum)
}

func visit(root *TreeNode, sum *int) *TreeNode {
	if nil == root {
		return root
	}
	visit(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	visit(root.Left, sum)
	return root
}
