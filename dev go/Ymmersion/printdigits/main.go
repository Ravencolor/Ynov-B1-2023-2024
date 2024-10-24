package main

import "github.com/01-edu/z01"

func main() {
	// Boucle pour parcourir les caractères ASCII des chiffres de 0 à 9
	for char := 48; char <= 57; char++ {
		// Affiche le caractère correspondant au code ASCII
		z01.PrintRune(rune(char))
	}
	// Affiche un retour à la ligne
	z01.PrintRune('\n')
}
