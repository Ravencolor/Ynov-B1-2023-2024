package piscine

// IsUpper vérifie si tous les caractères d'une chaîne sont des lettres majuscules.
func IsUpper(str string) bool {
	runeArray := []rune(str)     // Convertit la chaîne en un tableau de runes.
	runeCount := arrayCount(str) // Compte le nombre de runes dans la chaîne.
	count := 0
	for _, i := range runeArray {
		// Vérifie si la rune est une lettre majuscule.
		if i >= 'A' && i <= 'Z' {
			count++
		}
	}
	// Retourne vrai si toutes les runes sont des lettres majuscules.
	if count == runeCount {
		return true
	}
	return false
}
