package main

import "fmt"

func main() {
	var board [][]byte = [][]byte{
		[]byte{'A', 'B', 'C', 'D', 'E'},
		[]byte{'E', 'F', 'G', 'H', 'I'},
	}
	var word string = "ABCDEIHGFE"

	fmt.Println(exist(board, word))

}

func exist(board [][]byte, word string) bool {
	r, c, m := len(board), len(board[0]), len(word)
	var dfs func(int, int, int) bool
	dfs = func(i int, j int, idx int) bool {
		if i >= r ||
			j >= c ||
			i < 0 ||
			j < 0 ||
			idx >= m ||
			board[i][j] != word[idx] {
			return false
		}

		if idx == m-1 {
			return true
		}

		temp := board[i][j]
		board[i][j] = ' '

		ans := dfs(i+1, j, idx+1) ||
			dfs(i-1, j, idx+1) ||
			dfs(i, j+1, idx+1) ||
			dfs(i, j-1, idx+1)
		board[i][j] = temp

		return ans
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if dfs(i, j, 0) {
				return true
			}

		}
	}
	return false

}
