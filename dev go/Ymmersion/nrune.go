package piscine

// NRune retourne la nième rune d'une chaîne de caractères.
// Si n est hors des limites de la chaîne, la fonction retourne 0.
func NRune(s string, n int) rune {
	runes := []rune(s) // Convertit la chaîne en slice de runes
	len := 0
	for index := range runes {
		len = index // Calcule la longueur de la slice de runes
	}

	// Vérifie si n est dans les limites de la slice de runes
	if n-1 >= 0 && n-1 <= len {
		return runes[n-1] // Retourne la nième rune
	}
	return 0 // Retourne 0 si n est hors des limites
}
