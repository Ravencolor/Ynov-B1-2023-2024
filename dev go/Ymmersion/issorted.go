package piscine

// IsSorted vérifie si le tableau est trié selon la fonction de comparaison f.
// La fonction retourne true si le tableau est trié soit en ordre ascendant, soit en ordre descendant.
func IsSorted(f func(a, b int) int, tab []int) bool {
	ascendant, descendant := true, true

	// Vérifie si le tableau est trié en ordre descendant
	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) < 0 {
			descendant = false
		}
	}

	// Vérifie si le tableau est trié en ordre ascendant
	for i := 0; i < len(tab)-1; i++ {
		if f(tab[i], tab[i+1]) > 0 {
			ascendant = false
		}
	}

	// Retourne true si le tableau est trié soit en ordre ascendant, soit en ordre descendant
	return ascendant || descendant
}
