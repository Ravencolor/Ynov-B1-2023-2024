package piscine

// BasicAtoi convertit une chaîne de caractères en un entier.
// Si la chaîne contient des caractères non numériques, la conversion s'arrête et retourne la valeur actuelle.
func BasicAtoi(s string) int {
	arr := []rune(s) // Convertit la chaîne en une slice de runes
	n := len(s)      // Longueur de la chaîne
	v := 0           // Valeur entière résultante
	for i := 0; i < n; i++ {
		if arr[i] < '0' || arr[i] > '9' { // Vérifie si le caractère est un chiffre
			return v // Retourne la valeur actuelle si un caractère non numérique est trouvé
		} else {
			v *= 10                // Multiplie la valeur actuelle par 10 pour faire de la place pour le nouveau chiffre
			v += int(arr[i]) - '0' // Ajoute la valeur numérique du caractère à la valeur actuelle
		}
	}
	return v // Retourne la valeur entière convertie
}
