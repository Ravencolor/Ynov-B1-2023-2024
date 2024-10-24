package piscine

// SortWordArr trie un tableau de chaînes de caractères par ordre alphabétique
func SortWordArr(a []string) {
	// Boucle pour parcourir chaque élément du tableau sauf le dernier
	for i := 0; i < len(a)-1; i++ {
		// Boucle pour comparer l'élément actuel avec les éléments suivants
		for j := i + 1; j < len(a); j++ {
			// Si l'élément actuel est plus grand que l'élément suivant, on les échange
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
