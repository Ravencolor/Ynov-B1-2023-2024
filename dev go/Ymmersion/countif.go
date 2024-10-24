package piscine

// CountIf compte le nombre d'éléments dans le tableau arr pour lesquels la fonction f retourne true.
func CountIf(f func(string) bool, arr []string) int {
	compte := 0
	// Parcourt chaque élément du tableau arr
	for _, res := range arr {
		// Si la fonction f retourne true pour l'élément courant, incrémente le compteur
		if f(res) {
			compte += 1
		}
	}
	// Retourne le nombre d'éléments pour lesquels f a retourné true
	return compte
}
