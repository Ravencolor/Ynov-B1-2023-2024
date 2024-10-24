package main

import (
	"fmt"
	"os"
)

const N = 9
const Empty = '.'

// Vérifie si un numéro peut être placé dans une position donnée du tableau
func isValid(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < N; i++ {
		// Vérifie la ligne
		if board[row][i] == num {
			return false
		}
		// Vérifie la colonne
		if board[i][col] == num {
			return false
		}
		// Vérifie le sous-carré 3x3
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

// Résout le Sudoku en utilisant la récursion et le backtracking
func solveSudoku(board [][]byte) bool {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == Empty {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValid(board, i, j, num) {
						board[i][j] = num
						if solveSudoku(board) {
							return true
						}
						board[i][j] = Empty
					}
				}
				return false
			}
		}
	}
	return true
}

// Affiche le tableau Sudoku
func printBoard(board [][]byte) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// Vérifie si le nombre d'arguments est suffisant
	if len(os.Args) < 10 {
		fmt.Println("Error")
		return
	}
	board := make([][]byte, N)
	for i := 0; i < N; i++ {
		board[i] = make([]byte, N)
		for j := 0; j < N; j++ {
			arg := os.Args[i+1]
			// Vérifie si chaque argument a la longueur correcte
			if len(arg) != N {
				fmt.Println("Error")
				return
			}
			board[i][j] = arg[j]
		}
	}
	// Résout le Sudoku et affiche le résultat
	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
