package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Récupère les arguments de la ligne de commande, en ignorant le premier (le nom du programme)
	args := os.Args[1:]

	// Parcourt chaque argument
	for _, str := range args {
		// Parcourt chaque caractère de l'argument
		for _, ch := range str {
			// Affiche le caractère
			z01.PrintRune(ch)
		}
		// Affiche un saut de ligne après chaque argument
		z01.PrintRune('\n')
	}
}
