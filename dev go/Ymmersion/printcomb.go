package piscine

import "github.com/01-edu/z01"

// PrintComb affiche toutes les combinaisons possibles de trois chiffres différents
// dans l'ordre croissant, séparées par une virgule et un espace.
func PrintComb() {
	for i := '0'; i <= '7'; i++ { // Boucle pour le premier chiffre
		for j := i + 1; j <= '8'; j++ { // Boucle pour le deuxième chiffre
			for k := j + 1; k <= '9'; k++ { // Boucle pour le troisième chiffre
				z01.PrintRune(i) // Affiche le premier chiffre
				z01.PrintRune(j) // Affiche le deuxième chiffre
				z01.PrintRune(k) // Affiche le troisième chiffre

				// Vérifie si ce n'est pas la dernière combinaison pour ajouter une virgule et un espace
				if i != '7' || j != '8' || k != '9' {
					z01.PrintRune(',') // Affiche une virgule
					z01.PrintRune(' ') // Affiche un espace
				}
			}
		}
	}
	z01.PrintRune('\n') // Affiche un retour à la ligne à la fin
}
