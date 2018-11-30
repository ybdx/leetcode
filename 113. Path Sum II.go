package leetcode

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	result := make([]int, 0)
	getPath1(&res, root, result, sum)
	return res
}

func getPath1(res *[][]int, root *TreeNode, result []int, sum int) {
	temp := make([]int, len(result))
	copy(temp, result)
	if sum < 0 || nil == root {
		return
	}
	if nil == root.Left && nil == root.Right && root.Val == sum {
		result = append(result, root.Val)
		*res = append(*res, result)
	}
	temp = append(temp, root.Val)
	getPath1(res, root.Left, temp, sum - root.Val)
	getPath1(res, root.Right, temp, sum - root.Val)
}
