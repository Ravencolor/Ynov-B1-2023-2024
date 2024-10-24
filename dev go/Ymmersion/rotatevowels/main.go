package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Fonction pour vérifier si un caractère est une voyelle
func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}

func main() {
	// Récupère les arguments de la ligne de commande
	arg := os.Args[1:]
	rep := []rune{}
	ans := ""
	ln := 0
	IsF := true

	// Parcourt chaque argument
	for _, c := range arg {
		// Parcourt chaque caractère de l'argument
		for _, j := range c {
			// Si le caractère est une voyelle, l'ajoute à la liste rep
			if check(j) {
				rep = append(rep, j)
				ln++
			}
		}
		// Construit la chaîne de réponse ans
		if IsF {
			ans = c
			IsF = false
			continue
		}
		ans = ans + " " + c
	}

	cur := 0
	// Parcourt la chaîne de réponse ans
	for _, c := range ans {
		// Si le caractère est une voyelle, imprime la voyelle correspondante de la liste rep en ordre inverse
		if check(c) {
			z01.PrintRune(rep[ln-cur-1])
			cur++
		} else {
			// Sinon, imprime le caractère tel quel
			z01.PrintRune(c)
		}
	}
	// Imprime un saut de ligne à la fin
	z01.PrintRune('\n')
}
