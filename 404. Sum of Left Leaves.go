package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if nil == root {
		return sum
	}
	if nil != root && nil != root.Left && nil == root.Left.Left && nil == root.Left.Right {
		sum += root.Left.Val
	}
	sum += sumOfLeftLeaves(root.Left)
	sum += sumOfLeftLeaves(root.Right)
	return sum
}
