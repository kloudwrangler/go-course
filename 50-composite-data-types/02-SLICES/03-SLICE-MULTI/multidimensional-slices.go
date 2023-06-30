package main

import (
	"fmt"
)

func main() {
	// 1. Declare and initialize the tic-tac-toe board
	board := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}

	// 2. Print the board in a tabular format
	fmt.Println("Tic-Tac-Toe Board:")
	printBoard(board)

	// 3. Calculate and print the total number of X's and O's
	countX, countO := countXOs(board)
	fmt.Printf("Total X's: %d\nTotal O's: %d\n", countX, countO)

	// 4. Test the HasWinner() function
	hasWinner := hasWinner(board)
	fmt.Println("Has Winner?", hasWinner)

	// 5. Flip the board using the FlipBoard() function
	flipBoard(board)

	// 6. Print the flipped board
	fmt.Println("\nFlipped Board:")
	printBoard(board)
}

// printBoard prints the tic-tac-toe board in a tabular format
func printBoard(board [][]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

// countXOs counts the total number of X's and O's on the board
func countXOs(board [][]string) (int, int) {
	countX, countO := 0, 0
	for _, row := range board {
		for _, cell := range row {
			switch cell {
			case "X":
				countX++
			case "O":
				countO++
			}
		}
	}
	return countX, countO
}

// hasWinner checks if there is a winner on the board
func hasWinner(board [][]string) bool {
	for i := 0; i < 3; i++ {
		// Check rows
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
		// Check columns
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true
		}
	}
	// Check diagonals
	if (board[0][0] == board[1][1] && board[1][1] == board[2][2]) ||
		(board[0][2] == board[1][1] && board[1][1] == board[2][0]) {
		return true
	}
	return false
}

// flipBoard flips the values of X and O on the board
func flipBoard(board [][]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "X" {
				board[i][j] = "O"
			} else if board[i][j] == "O" {
				board[i][j] = "X"
			}
		}
	}
}
