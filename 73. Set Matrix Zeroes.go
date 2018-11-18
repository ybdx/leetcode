package leetcode

func setZeroes(matrix [][]int) {
	isCol := false // 当第一列有0存在时为true
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		if 0 == matrix[i][0] {
			isCol = true
		}
		for j := 1; j < n; j++ {
			if 0 == matrix[i][j] {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if 0 == matrix[i][0] || 0 == matrix[0][j] {
				matrix[i][j] = 0
			}
		}
	}
	if 0 == matrix[0][0] {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if isCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
