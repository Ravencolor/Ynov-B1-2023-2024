package piscine

// BasicAtoi2 convertit une chaîne de caractères en un entier.
// Si la chaîne contient des caractères non numériques, elle retourne 0.
func BasicAtoi2(s string) int {
	arr := []rune(s) // Convertit la chaîne en une slice de runes
	n := len(s)      // Longueur de la chaîne
	v := 0           // Valeur entière résultante
	for i := 0; i < n; i++ {
		// Vérifie si le caractère est un chiffre
		if arr[i] < '0' || arr[i] > '9' {
			return 0 // Retourne 0 si un caractère non numérique est trouvé
		} else {
			v *= 10                // Décale la valeur actuelle d'un ordre de grandeur
			v += int(arr[i]) - '0' // Ajoute la valeur numérique du caractère
		}
	}
	return v // Retourne la valeur entière convertie
}
