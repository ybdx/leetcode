package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNode2 struct {
	Val int
	Child []*TreeNode2
}
