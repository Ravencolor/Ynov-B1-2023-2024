package piscine

import "github.com/01-edu/z01"

// PrintStr affiche chaque caractère de la chaîne passée en argument
func PrintStr(s string) {
	// Boucle sur chaque caractère de la chaîne
	for _, char := range s {
		// Affiche le caractère courant
		z01.PrintRune(char)
	}
}
