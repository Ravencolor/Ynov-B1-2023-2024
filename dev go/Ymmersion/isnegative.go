package piscine

import "github.com/01-edu/z01"

// IsNegative vérifie si un nombre est négatif ou non.
// Si le nombre est négatif, il affiche 'T'.
// Sinon, il affiche 'F'.
func IsNegative(nb int) {
	if nb >= 0 {
		// Si le nombre est positif ou nul, afficher 'F'
		z01.PrintRune('F')
		z01.PrintRune('\n')
	} else {
		// Si le nombre est négatif, afficher 'T'
		z01.PrintRune('T')
		z01.PrintRune('\n')
	}
}
