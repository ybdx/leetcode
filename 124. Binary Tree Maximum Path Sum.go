package leetcode

func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 32
	computerMaxSum(&maxSum, root)
	return maxSum
}

func computerMaxSum(maxSum *int, root *TreeNode) int {
	if nil == root {
		return 0
	}
	leftMax := max(0, computerMaxSum(maxSum, root.Left))
	rightMax := max(0, computerMaxSum(maxSum, root.Right))
	*maxSum = max(*maxSum, leftMax + rightMax + root.Val)
	return max(leftMax, rightMax) + root.Val
}
