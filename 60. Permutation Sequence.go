package leetcode

import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	source := make([]int, n)
	factor := make([]int, n+1)
	res := make([]string, 0)
	for i:=0; i<n; i++ {
		source[i] = i+1
	}
	factor[0] = 1
	for i:=1; i<=n; i++ {
		factor[i] = i * factor[i-1]
	}
	k--
	for i:=1; i<=n; i++ {
		index := k / factor[n-i]
		res = append(res, strconv.Itoa(source[index]))
		source = append(source[:index], source[index+1:]...)
		k %= factor[n-i]
	}
	return strings.Join(res, "")
}
