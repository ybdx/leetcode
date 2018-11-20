package leetcode

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})
	size := 0
	for i := 0; i < len(nums); i++ {
		size = len(result)
		for j:=0; j<size; j++ {
			temp := make([]int, len(result[j]))
			copy(temp, result[j])
			temp = append(temp, nums[i])
			result = append(result, temp)
		}
	}
	return result
}
