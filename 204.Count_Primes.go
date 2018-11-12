package leetcode

func countPrimes(n int) int {
	if 2 > n {
		return 0
	}
	count := 0
	flag := make([]bool, n)
	for i := 2; i < n; i++ {
		if !flag[i] {
			count++
		}
		if flag[i] {
			continue
		}
		for j := 2; i*j < n; j++ {
			flag[i*j] = true
		}
	}
	return count
}
