package piscine

// TrimAtoi convertit une chaîne de caractères en un entier, en ignorant les caractères non numériques.
// Si la chaîne contient un signe négatif avant les chiffres, le résultat sera négatif.
func TrimAtoi(s string) int {
	var neg bool = false  // Indique si le nombre est négatif
	var empty bool = true // Indique si aucun chiffre n'a été trouvé
	var res int = 0       // Résultat final

	for _, v := range s {
		if empty && !neg && v == '-' {
			neg = true // Marque le nombre comme négatif si un '-' est trouvé avant les chiffres
		} else if IsRuneDigit(v) {
			res *= 10          // Décale les chiffres vers la gauche
			res += int(v - 48) // Ajoute le chiffre actuel au résultat
			empty = false      // Indique qu'au moins un chiffre a été trouvé
		}
	}

	if empty {
		return 0 // Retourne 0 si aucun chiffre n'a été trouvé
	} else {
		if neg {
			return -res // Retourne le résultat négatif si un '-' a été trouvé
		} else {
			return res // Retourne le résultat positif
		}
	}
}

// IsRuneDigit vérifie si une rune est un chiffre.
func IsRuneDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true // Retourne vrai si la rune est un chiffre
	}
	return false // Retourne faux sinon
}
