package piscine

// Join prend un tableau de chaînes de caractères et un séparateur,
// et retourne une seule chaîne de caractères où chaque élément du tableau
// est séparé par le séparateur.
func Join(strs []string, sep string) string {
	strV := ""
	// Boucle à travers chaque élément du tableau de chaînes de caractères.
	for i := 0; i < len(strs); i++ {
		strV += strs[i] // Ajoute l'élément actuel à la chaîne de résultat.
		// Ajoute le séparateur si ce n'est pas le dernier élément.
		if i < len(strs)-1 {
			strV += sep
		}
	}
	return strV // Retourne la chaîne de caractères résultante.
}
