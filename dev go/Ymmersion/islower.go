package piscine

// IsLower vérifie si tous les caractères de la chaîne sont en minuscules
func IsLower(str string) bool {
	runeArray := []rune(str)     // Convertit la chaîne en un tableau de runes
	runeCount := arrayCount(str) // Compte le nombre de runes dans la chaîne
	count := 0
	for _, i := range runeArray {
		// Vérifie si le caractère est une lettre minuscule
		if i >= 'a' && i <= 'z' {
			count++
		}
	}
	// Retourne true si tous les caractères sont en minuscules, sinon false
	if count == runeCount {
		return true
	}
	return false
}
