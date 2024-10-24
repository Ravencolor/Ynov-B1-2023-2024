package main

import (
	"os"

	"github.com/01-edu/z01"
)

// SortTable trie un tableau de chaînes de caractères par ordre alphabétique
func SortTable(table []string) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func main() {
	// Récupère les arguments de la ligne de commande, en ignorant le premier argument (le nom du programme)
	args := os.Args[1:]

	// Trie les arguments
	SortTable(args)

	// Affiche les arguments triés
	for _, arg := range args {
		for _, b := range arg {
			z01.PrintRune(rune(b))
		}
		z01.PrintRune('\n')
	}
}
