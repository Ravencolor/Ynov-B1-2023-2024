package piscine

// Map applique la fonction f à chaque élément du tableau arr
// et retourne un nouveau tableau contenant les résultats.
func Map(f func(int) bool, arr []int) []bool {
	// Crée un nouveau tableau de booléens de la même longueur que arr.
	fin := make([]bool, len(arr))

	// Parcourt chaque élément de arr, applique la fonction f,
	// et stocke le résultat dans le tableau fin.
	for i, res := range arr {
		fin[i] = f(res)
	}

	// Retourne le tableau contenant les résultats.
	return fin
}
