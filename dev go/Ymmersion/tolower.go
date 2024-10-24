package piscine

// ToLower convertit tous les caractères d'une chaîne en minuscules
func ToLower(s string) string {
	h := []rune(s) // Convertit la chaîne en slice de runes pour gérer les caractères Unicode
	result := ""   // Initialise la chaîne de résultat
	for i := 0; i <= len(h)-1; i++ {
		// Vérifie si le caractère est une majuscule
		if (h[i] >= 'A') && (h[i] <= 'Z') {
			h[i] = h[i] + 32 // Convertit la majuscule en minuscule
		}
		result += string(h[i]) // Ajoute le caractère (converti ou non) au résultat
	}
	return result // Retourne la chaîne convertie en minuscules
}
