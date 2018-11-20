package leetcode

import (
	"sort"
)

// eg.[1, 2, 2]
// 初始 [[]] startIndex = 0 size = 0
// 1	[[], [1]] startIndex = 0 size = 1
// 2	[[], [1], [2], [1, 2]] startIndex = 0 size = 2
// 2	[[], [1], [2], [1, 2], [2, 2], [1, 2, 2]] startIndex = 2 size = 4
func subsetsWithDup(nums []int) [][]int {
	sliceNums := sort.IntSlice(nums)
	sort.Sort(sliceNums)
	startIndex, size := 0, 0 // startIndex 表示前面结果的长度
	result := make([][]int, 0)
	result = append(result, make([]int, 0))
	for i := 0; i < len(nums); i++ {
		if i > 0 && sliceNums[i] == sliceNums[i-1] {
			startIndex = size
		}
		size = len(result)
		for j := startIndex; j < size; j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			temp = append(temp, sliceNums[i])
			result = append(result, temp)
		}
	}
	return result
}
