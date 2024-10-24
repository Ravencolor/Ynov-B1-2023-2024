package main

import "github.com/01-edu/z01"

func main() {
	// Boucle pour afficher les lettres de l'alphabet en ordre inverse
	for char := (97 + 25); char >= (97); char-- {
		z01.PrintRune(rune(char)) // Affiche la lettre courante
	}
	z01.PrintRune('\n') // Affiche un retour à la ligne
}
