package piscine

// Sqrt calcule la racine carrée entière d'un nombre donné.
// Si le nombre n'a pas de racine carrée entière, la fonction retourne 0.
func Sqrt(nb int) int {
	// Si le nombre est 1, la racine carrée est 1.
	if nb == 1 {
		return 1
	}
	// Si le nombre est 2, il n'a pas de racine carrée entière.
	if nb == 2 {
		return 0
	}
	// Si le nombre est positif.
	if nb > 0 {
		result := 1
		sqrt := 0
		// Boucle pour trouver la racine carrée entière.
		for a := 1; result <= nb; a++ {
			result = a * a
			sqrt++
			// Si le carré de 'a' est égal au nombre, retourner 'sqrt'.
			if result == nb {
				return sqrt
			}
		}
		// Si aucune racine carrée entière n'a été trouvée, retourner 0.
		return 0
	} else {
		// Si le nombre est négatif, retourner 0.
		return 0
	}
}
