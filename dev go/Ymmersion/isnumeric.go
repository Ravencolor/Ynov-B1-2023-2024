package piscine

// IsNumeric vérifie si une chaîne de caractères est composée uniquement de chiffres.
func IsNumeric(str string) bool {
	h := []rune(str)
	// Parcourt chaque caractère de la chaîne.
	for i := 0; i <= len(h)-1; i++ {
		// Vérifie si le caractère n'est pas un chiffre.
		if (h[i] >= 0) && (h[i] <= 47) || (h[i] >= 58) && (h[i] <= 127) {
			return false
		}
	}
	// Retourne vrai si tous les caractères sont des chiffres.
	return true
}
