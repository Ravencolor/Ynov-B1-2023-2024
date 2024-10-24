package piscine

// AlphaCount compte le nombre de caractères alphabétiques dans une chaîne de caractères.
func AlphaCount(str string) int {
	count := 0
	for _, s := range str {
		if isAlpha(s) {
			count++
		}
	}
	return count
}

// isAlpha vérifie si un caractère est alphabétique (majuscule ou minuscule).
func isAlpha(alpha rune) bool {
	// Vérifie les lettres minuscules
	for a := 'a'; a <= 'z'; a++ {
		if alpha == a {
			return true
		}
	}
	// Vérifie les lettres majuscules
	for a := 'A'; a <= 'Z'; a++ {
		if alpha == a {
			return true
		}
	}
	return false
}
