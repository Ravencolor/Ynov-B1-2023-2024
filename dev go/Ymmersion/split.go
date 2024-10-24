package piscine

// Split divise une chaîne de caractères 'str' en utilisant un 'charset' comme séparateur
func Split(str, charset string) []string {
	ln := 0
	// Calculer la longueur du charset
	for i := range charset {
		ln = i + 1
	}
	ln2 := 0
	// Calculer la longueur de la chaîne str
	for i := range str {
		ln2 = i + 1
	}
	// Remplacer chaque occurrence de charset dans str par un espace
	for i := 0; i < ln2-ln; i++ {
		if str[i:i+ln] == charset {
			str = str[:i] + " " + str[i+ln:]
			ln2 -= ln
		}
	}
	// Appeler la fonction SplitWhiteSpaces pour diviser la chaîne par les espaces
	return SplitWhiteSpaces(str)
}

// lent2 calcule la longueur d'une chaîne de caractères
func lent2(d string) int {
	inc := 0
	// Incrémenter pour chaque caractère dans la chaîne
	for range d {
		inc++
	}
	return inc
}
