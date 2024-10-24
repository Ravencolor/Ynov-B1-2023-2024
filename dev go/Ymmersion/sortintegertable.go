package piscine

// SortIntegerTable trie un tableau d'entiers en utilisant l'algorithme de tri à bulles.
func SortIntegerTable(table []int) {
	n := len(table)
	// Parcourt tous les éléments du tableau
	for i := 0; i < n-1; i++ {
		// Parcourt le tableau de 0 à n-i-1
		for j := 0; j < n-i-1; j++ {
			// Échange les éléments si l'élément actuel est plus grand que le suivant
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
