package leetcode

func getSum(a int, b int) int {
	if 0 == b {
		return a
	}
	if 0 == a {
		return b
	}
	for b != 0 {
		temp := a & b
		a = a ^ b
		b = temp << 1
	}
	return a
}
