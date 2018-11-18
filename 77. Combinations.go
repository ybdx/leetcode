package leetcode

func combine(n int, k int) [][]int {
	if k < 0 || k > n {
		return [][]int{}
	}
	res := make([][]int, 0)
	temp := make([]int, 0)
	genResult(&res, &temp, 1, n, k)
	return res
}

func genResult(res *[][]int, temp *[]int, start, end, k int) {
	if len(*temp) == k {
		dest := make([]int, k)
		copy(dest, *temp)
		*res = append(*res, dest)
	}
	for i := start; i <= end; i++ {
		*temp = append(*temp, i)
		genResult(res, temp, i+1, end, k)
		*temp = (*temp)[:len(*temp)-1]
	}
}
