package leetcode

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	flag := 0
	result := make([]string, 0)
	for i >= 0 || j >= 0 || flag == 1 {
		temp1 := 0
		temp2 := 0
		if i >= 0 {
			temp1 = int(num1[i] - '0')
		}
		if j >= 0 {
			temp2 = int(num2[j] - '0')
		}
		temp := temp1 + temp2 + flag
		result = append(result,  strconv.Itoa(temp % 10))
		flag = temp / 10
		i--
		j--
	}
	for i, j = 0, len(result) - 1; i < j; i,j=i+1,j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return strings.Join(result, "")
}
