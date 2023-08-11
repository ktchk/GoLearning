package main

import (
	"strconv"
)

/*

-------------------------
| 7 0 8 | 0 0 1 | 4 0 0 |
| 3 0 0 | 4 7 0 | 0 0 0 |
| 0 9 4 | 5 0 0 | 0 7 6 |
-------------------------
| 2 7 0 | 6 0 8 | 0 0 0 |
| 6 0 3 | 0 0 0 | 7 0 2 |
| 0 0 0 | 7 0 0 | 1 6 0 |
-------------------------
| 0 0 7 | 0 5 4 | 0 0 1 |
| 4 1 5 | 0 0 0 | 8 2 7 |
| 0 0 0 | 0 0 7 | 0 0 4 |
-------------------------

*/

/* Plan:
1. [][]int to store the puzzle
2. display the puzzle func
3. func that will find index of 0 and if not = puzzle is solved
4. func than will check the possibility of entering (1-9) in that index in table
5. ...
6. profit!

*/

func displayTheBoard(board [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			print(strconv.Itoa(board[i][j]) + " ")
		}
		println()
	}
}

func findZeroIndex(board [][]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isNumValid(board [][]int, guess, row, column int) bool {
	// 1. check in row
	// 2. check in column
	// 3. check in 3x3 square

	for index := range board {
		if board[row][index] == guess && column != index {
			return false
		}
	}

	for index := range board {
		if board[index][column] == guess && row != index {
			return false
		}
	}

	// Row - (Row % 3) + Value for cycling

	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if board[row-row%3+k][column-column%3+l] == guess && (row-row%3+k != row || column-column%3+l != column) {
				return false
			}
		}

	}
	return true
}

func solvePuzzle(board [][]int, count int) bool {
	row, column := findZeroIndex(board)
	if row == -1 {
		println("The puzzle is solved!")
		println("number of iterations:", count)
		return true
	} else {
		row, column = findZeroIndex(board)
	}
	for i := 1; i <= 9; i++ {
		count++
		if isNumValid(board, i, row, column) {
			board[row][column] = i

			if solvePuzzle(board, count) {
				return true
			}
			board[row][column] = 0
		}
	}
	return false
}

func main() {
	board := [][]int{
		{0, 9, 4, 1, 0, 0, 0, 0, 0},
		{3, 0, 0, 0, 0, 6, 4, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
		{2, 0, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 1, 6, 0, 0, 8, 2, 0},
		{0, 0, 0, 0, 0, 8, 0, 0, 3},
		{0, 0, 3, 0, 5, 0, 0, 0, 0},
		{5, 0, 0, 0, 0, 4, 2, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 0},
	}

	displayTheBoard(board)
	count := 0
	solvePuzzle(board, count)
	displayTheBoard(board)
}
