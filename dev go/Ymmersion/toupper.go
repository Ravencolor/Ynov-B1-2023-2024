package piscine

// ToUpper convertit une chaîne de caractères en majuscules
func ToUpper(s string) string {
	h := []rune(s)                   // Convertit la chaîne en slice de runes pour gérer les caractères Unicode
	result := ""                     // Initialise la chaîne de résultat
	for i := 0; i <= len(h)-1; i++ { // Parcourt chaque rune de la chaîne
		if (h[i] >= 'a') && (h[i] <= 'z') { // Vérifie si la rune est une lettre minuscule
			h[i] = h[i] - 32 // Convertit la lettre minuscule en majuscule
		}
		result += string(h[i]) // Ajoute la rune (convertie ou non) à la chaîne de résultat
	}
	return result // Retourne la chaîne convertie en majuscules
}
