package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return root
	}
	right := invertTree(root.Left)
	left := invertTree(root.Right)
	root.Left = left
	root.Right = right
	return root
}
