package leetcode

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	return getSum(root, &sum)
}

func getSum(root *TreeNode, sum *int) *TreeNode {
	if nil == root {
		return root
	}
	getSum(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	getSum(root.Left, sum)
	return root
}
