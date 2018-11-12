package leetcode

func isUgly(num int) bool {
	prime := []int{2, 3, 5}
	if 0 == num {
		return false
	}
	for _, v := range prime {
		for num%v == 0 {
			num /= v
		}
	}
	if 1 == num {
		return true
	}
	return false
}
