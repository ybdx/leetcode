package leetcode

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	m := len(dungeon)
	n := len(dungeon[0])
	hp := make([][]int, 0)
	for i:=0; i< m;i++ {
		temp := make([]int, n)
		hp = append(hp, temp)
	}
	hp[m - 1][n - 1] = max(-dungeon[m-1][n-1] + 1, 1)
	for i:=m-1; i>0; i-- {
		hp[i-1][n-1] = max(hp[i][n-1] - dungeon[i-1][n-1], 1)
	}
	for j :=n-1; j>0; j-- {
		hp[m-1][j-1] = max(hp[m-1][j] - dungeon[m-1][j-1], 1)
	}

	for i:=m-2;i >=0; i-- {
		for j:=n-2;j>=0; j-- {
			right := max(hp[i+1][j] - dungeon[i][j], 1)
			down := max(hp[i][j+1] - dungeon[i][j], 1)
			hp[i][j] = min(right, down)
		}
	}

	return hp[0][0]
}