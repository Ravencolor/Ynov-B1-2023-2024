package piscine

// IsAlpha vérifie si une chaîne de caractères ne contient que des lettres alphabétiques.
func IsAlpha(str string) bool {
	h := []rune(str)
	// Parcourt chaque caractère de la chaîne.
	for i := 0; i <= len(h)-1; i++ {
		// Vérifie si le caractère n'est pas une lettre alphabétique.
		if (h[i] >= 0) && (h[i] <= 47) || (h[i] >= 58) && (h[i] <= 64) || (h[i] >= 91) && (h[i] <= 96) || (h[i] >= 123) && (h[i] <= 127) {
			return false
		}
	}
	// Retourne vrai si tous les caractères sont alphabétiques.
	return true
}
