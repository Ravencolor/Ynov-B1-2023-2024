package piscine

import "github.com/01-edu/z01"

// PrintNbrBase convertit un nombre entier en une chaîne de caractères
// dans une base donnée et l'affiche.
func PrintNbrBase(s int, t string) {
	ans := ""
	ln := 0

	// Calculer la longueur de la base
	for _, c := range t {
		if c == c {
			ln++
		}
	}
	mx_p := ln

	// Gérer les nombres négatifs
	if s < 0 {
		ans = "-"
		mx_p *= -1
	}

	// Vérifier si la base est valide (longueur > 1)
	if ln > 1 {

		// Calculer la puissance maximale de la base
		for s/mx_p >= ln {
			mx_p *= ln
		}

		// Convertir le nombre en base donnée
		for mx_p != 0 {
			ans = ans + string(t[s/mx_p])
			s = s - s/mx_p*mx_p
			mx_p /= ln
		}

		// Vérifier les caractères invalides dans la base
		x := map[rune]bool{}
		for _, c := range t {
			if c == '+' || c == '-' {
				ans = "NV"
				break
			}
			if x[c] == false {
				x[c] = true
			} else {
				ans = "NV"
				break
			}
		}
	} else {
		ans = "NV"
	}

	// Afficher le résultat
	for _, c := range ans {
		z01.PrintRune(c)
	}
}
