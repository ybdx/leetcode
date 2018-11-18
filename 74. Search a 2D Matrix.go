package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	for i < m {
		if matrix[i][n -1] == target {
			return true
		} else if matrix[i][n -1] > target {
			break
		} else {
			i++
		}
	}
	if i == m {
		return false
	}
	return binarySearch(matrix[i], 0, n-1, target)
}

func binarySearch(nums []int, start, end, target int) bool {
	for start <= end {
		mid := start + (end - start) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			end = mid -1
		} else {
			start = mid + 1
		}
	}
	return false
}
