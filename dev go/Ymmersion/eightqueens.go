package piscine

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

// Vérifie si la position de la reine est sûre
func isSafe(queen_number, row_position int) bool {
	for i := 0; i < queen_number; i++ {
		other_row_pos := position[i]

		// Vérifie les conflits avec les autres reines
		if other_row_pos == row_position || other_row_pos == row_position-(queen_number-i) || other_row_pos == row_position+(queen_number-i) {
			return false
		}
	}
	return true
}

// Résout le problème des huit reines
func solve(k int) {
	if k == N {
		// Affiche la solution
		for i := 0; i < N; i++ {
			z01.PrintRune(rune(position[i] + '1'))
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				solve(k + 1)
			}
		}
	}
}

// Fonction principale pour lancer la résolution
func EightQueens() {
	solve(0)
}
