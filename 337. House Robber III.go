package leetcode

func rob1(root *TreeNode) int {
	return max(robMoney(root, true), robMoney(root, false))
}

func robMoney(root *TreeNode, visit bool) int {
	if nil == root {
		return 0
	}
	if visit {
		return  robMoney(root.Left, false) + robMoney(root.Right, false)
	}
	return max(root.Val + robMoney(root.Left, true) + robMoney(root.Right, true), robMoney(root.Left, false) + robMoney(root.Right, false))
}

