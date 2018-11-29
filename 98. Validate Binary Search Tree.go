package leetcode

func isValidBST(root *TreeNode) bool {
	return checkBST(root, -1 * 1 << 32, 1 << 32)
}

func checkBST(root *TreeNode, min, max int) bool {
	if nil == root {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return checkBST(root.Left, min, root.Val) && checkBST(root.Right, root.Val, max)
}
