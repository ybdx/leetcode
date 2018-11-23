package leetcode

// f(n) = f(n-1)*f(0) + f(n-2)*f(1) + ... + f(n-i)*f(i-1) + ... + f(0)*f(n-1)
// f(n) = C(2n,n)/(n+1)
// f(n)=f(n-1)*(4*n-2)/(n+1);
// 上面该公式是卡特兰数, 相关推导公司见wiki: https://blog.csdn.net/hemeinvyiqiluoben/article/details/11320419
func numTrees(n int) int {
	if 0 == n || 1 == n {
		return 1
	}
	res := 1
	for i:=2; i <=n; i++ {
		res = res * (4 * i - 2) / (i + 1)
	}
	return res
}

//func numTrees(n int) int {
//	if 0 == n || 1 == n {
//		return 1
//	}
//	G := make([]int, 0)
//	G = append(G, 1)
//	G = append(G, 1)
//	for i:=2; i<=n; i++ {
//		temp := 0
//		for j:=1; j <=i; j++ {
//			temp += G[j-1] * G[i-j]
//		}
//		G = append(G, temp)
//	}
//	return G[n]
//}