package leetcode

// 二进制生成格雷码方法： 二进制第i位和第i+1位相同，则格雷码第i位为0，对应的就是 i ^ i >> 1
func grayCode(n int) []int {
	res := make([]int, 0)
	for i:=0; i< 1<<uint(n); i++ {
		res = append(res, i ^ i >> 1)
	}
	return res
}
