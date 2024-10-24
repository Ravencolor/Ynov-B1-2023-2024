package piscine

import "github.com/01-edu/z01"

// PrintWordsTables prend un tableau de chaînes de caractères et imprime chaque chaîne sur une nouvelle ligne
func PrintWordsTables(table []string) {
	// Parcourt chaque chaîne dans le tableau
	for _, c := range table {
		myStr := c
		// Parcourt chaque rune dans la chaîne et l'imprime
		for _, j := range myStr {
			z01.PrintRune(j)
		}
		// Imprime un saut de ligne après chaque chaîne
		z01.PrintRune('\n')
	}
}
