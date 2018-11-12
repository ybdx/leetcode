package leetcode

import (
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	str := make([]string, len(nums))
	// 将数字转化为字符串
	for k, v := range nums {
		str[k] = strconv.Itoa(v)
	}
	quickSort(str)
	if str[0] == "0" {
		return "0"
	}
	result := strings.Join(str, "")
	return result
}

func quickSort(nums []string) {
	if len(nums) <= 1 {
		return
	}
	mid, i := nums[0], 1
	head, tail := 0, len(nums) - 1
	for head < tail {
		if compare(nums[i], mid) {
			nums[i], nums[head] = nums[head], nums[i]
			head ++
			i++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail --
		}
	}
	nums[head] = mid
	quickSort(nums[:head])
	quickSort(nums[head+1:])
}

// 判断字符串大小
func compare(a, b string) bool {
	order1 := a + b
	order2 := b + a
	if order1 > order2 {
		return true
	}
	return false
}
