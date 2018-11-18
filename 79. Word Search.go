package leetcode

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	if rows == 0 || cols == 0 {
		return false
	}
	var visitor [][]bool
	for i := 0; i < rows; i++ {
		temp := make([]bool, cols)
		visitor = append(visitor, temp)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backup(board, word, 0, i, j, &visitor) {
				return true
			}
		}
	}
	return false
}

func backup(board [][]byte, word string, index, start, end int, visitor *[][]bool) bool {
	if index == len(word) {
		return true
	}
	if start < 0 || start == len(board) || end < 0 || end == len(board[0]) {
		return false
	}
	if word[index] != board[start][end] {
		return false
	}
	if (*visitor)[start][end] {
		return false
	}
	(*visitor)[start][end] = true
	exist := backup(board, word, index+1, start+1, end, visitor) ||
		backup(board, word, index+1, start, end+1, visitor) ||
		backup(board, word, index+1, start-1, end, visitor) ||
		backup(board, word, index+1, start, end-1, visitor)
	if !exist {
		(*visitor)[start][end] = false
	}
	return exist
}
