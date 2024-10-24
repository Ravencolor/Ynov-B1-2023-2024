package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Récupère les arguments de la ligne de commande, en ignorant le premier argument (le nom du programme)
	args := os.Args[1:]
	// Initialise l'index à la fin de la liste des arguments
	args_index := len(args) - 1
	// Boucle sur les arguments en ordre inverse
	for args_index >= 0 {
		// Boucle sur chaque caractère de l'argument courant
		for _, str := range args[args_index] {
			// Affiche le caractère courant
			z01.PrintRune(str)
		}
		// Décrémente l'index pour passer à l'argument précédent
		args_index--
		// Affiche un saut de ligne après chaque argument
		z01.PrintRune('\n')
	}
}
