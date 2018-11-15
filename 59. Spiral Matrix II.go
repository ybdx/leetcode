package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		temp := make([]int, n)
		res = append(res, temp)
	}
	curNum := 0
	for start := 0; start <= n/2; start++ {
		getMatrixCircle(res, n, n, start, &curNum)
	}
	return res
}

func getMatrixCircle(mat [][]int, row, col, start int, curNum *int) {
	endX := col - start - 1
	endY := row - start - 1
	// 从左到右
	for i := start; i <= endX; i++ {
		*curNum ++
		mat[start][i] = *curNum
	}
	// 从上到下
	if start < endY {
		for i := start + 1; i <= endY; i++ {
			*curNum++
			mat[i][endX] = *curNum
		}
	}
	// 从右到左
	if start < endX && start < endY {
		for i := endX - 1; i >= start; i-- {
			*curNum++
			mat[endY][i] = *curNum
		}
	}
	// 从下到上
	if endY > start+1 && start < endX {
		for i := endY - 1; i > start; i-- {
			*curNum++
			mat[i][start] = *curNum
		}
	}
}
