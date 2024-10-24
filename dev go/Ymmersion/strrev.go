package piscine

// StrRev prend une chaîne de caractères en entrée et retourne la chaîne inversée.
func StrRev(s string) string {
	// Convertit la chaîne en un tableau de runes pour gérer correctement les caractères Unicode.
	chars := []rune(s)
	// Boucle pour inverser les runes dans le tableau.
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		// Échange les runes aux positions i et j.
		chars[i], chars[j] = chars[j], chars[i]
	}
	// Convertit le tableau de runes inversé en une chaîne et la retourne.
	return string(chars)
}
