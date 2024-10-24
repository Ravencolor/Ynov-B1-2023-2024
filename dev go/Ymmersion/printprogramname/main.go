package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Récupère le nom du programme (chemin complet)
	name := os.Args[0]
	// Parcourt chaque caractère du nom du programme à partir du 3ème caractère
	for _, ch := range name[2:] {
		// Affiche chaque caractère
		z01.PrintRune(ch)
	}
	// Affiche un retour à la ligne
	z01.PrintRune('\n')
}
