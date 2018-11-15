package leetcode

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}
	start := 0
	fmt.Println(len(matrix), len(matrix[0]))
	for start * 2 < len(matrix) && start * 2 < len(matrix[0]) {
		getSpiralOrder(matrix, len(matrix), len(matrix[0]), start, &res)
		start++
	}
	return res
}

func getSpiralOrder(matrix [][]int, row, col, start int, res *[]int) {
	endX := col - start - 1
	endY := row - start - 1
	// 左->右
	for i := start; i <= endX; i++ {
		*res = append(*res, matrix[start][i])
	}
	// 上->下
	if endY > start {
		for i := start + 1; i <= endY; i++ {
			*res = append(*res, matrix[i][endX])
		}
	}
	// 右->左
	if endX > start && endY > start {
		for i := endX - 1; i >= start; i-- {
			*res = append(*res, matrix[endY][i])
		}
	}
	// 下到上
	if endY > start+1 && endX > start {
		for i := endY - 1; i > start; i-- {
			*res = append(*res, matrix[i][start])
		}
	}
}
