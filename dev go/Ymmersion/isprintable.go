package piscine

// IsPrintable vérifie si tous les caractères de la chaîne sont imprimables
func IsPrintable(str string) bool {
	h := []rune(str)
	// Parcourt chaque caractère de la chaîne
	for i := 0; i <= len(h)-1; i++ {
		// Vérifie si le caractère est un caractère de contrôle non imprimable
		if (h[i] >= 0) && (h[i] <= 31) {
			return false
		}
	}
	return true
}
