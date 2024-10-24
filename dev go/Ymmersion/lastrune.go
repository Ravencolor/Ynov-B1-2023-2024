package piscine

// LastRune retourne la dernière rune d'une chaîne de caractères.
func LastRune(str string) rune {
	// Convertit la chaîne de caractères en slice de runes.
	lastrune := []rune(str)
	// Calcule la longueur du slice de runes.
	lens := len(lastrune)
	// Retourne la dernière rune du slice.
	return lastrune[lens-1]
}
