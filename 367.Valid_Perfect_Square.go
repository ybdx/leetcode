package leetcode

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	if num < 2 {
		return true
	}
	start := 2
	end := num >> 1
	for start <= end {
		mid := (start + end) >> 1
		temp := mid * mid
		if temp == num {
			return true
		} else if temp > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
