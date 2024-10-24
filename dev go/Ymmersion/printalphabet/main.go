package main

import "github.com/01-edu/z01"

func main() {
	// Boucle pour parcourir les lettres de l'alphabet en minuscule (de 'a' à 'z')
	for char := 97; char <= (97 + 25); char++ {
		// Affiche chaque lettre
		z01.PrintRune(rune(char))
	}
	// Affiche un retour à la ligne à la fin
	z01.PrintRune('\n')
}
