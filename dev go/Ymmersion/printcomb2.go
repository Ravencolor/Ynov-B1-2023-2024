package piscine

import "github.com/01-edu/z01"

// PrintComb2 affiche toutes les combinaisons de deux nombres à deux chiffres
func PrintComb2() {
	for a := '0'; a <= '9'; a = a + 1 { // Boucle pour le premier chiffre du premier nombre
		for b := '0'; b <= '9'; b = b + 1 { // Boucle pour le deuxième chiffre du premier nombre
			d := b + 1                        // Initialisation du deuxième chiffre du deuxième nombre
			for c := a; c <= '9'; c = c + 1 { // Boucle pour le premier chiffre du deuxième nombre
				for ; d <= '9'; d = d + 1 { // Boucle pour le deuxième chiffre du deuxième nombre
					z01.PrintRune(a)                              // Affiche le premier chiffre du premier nombre
					z01.PrintRune(b)                              // Affiche le deuxième chiffre du premier nombre
					z01.PrintRune(' ')                            // Affiche un espace
					z01.PrintRune(c)                              // Affiche le premier chiffre du deuxième nombre
					z01.PrintRune(d)                              // Affiche le deuxième chiffre du deuxième nombre
					if a < '9' || b < '8' || c < '9' || d < '9' { // Vérifie si ce n'est pas la dernière combinaison
						z01.PrintRune(',') // Affiche une virgule
						z01.PrintRune(' ') // Affiche un espace
					}
				}
				d = '0' // Réinitialise le deuxième chiffre du deuxième nombre
			}
		}
	}
	z01.PrintRune('\n') // Affiche un retour à la ligne à la fin
}
