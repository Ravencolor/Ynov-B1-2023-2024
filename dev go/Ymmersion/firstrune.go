package piscine

// FirstRune retourne la première rune d'une chaîne de caractères
func FirstRune(str string) rune {
	// Convertit la chaîne de caractères en slice de runes
	firstrune := []rune(str)
	// Retourne la première rune
	return firstrune[0]
}
