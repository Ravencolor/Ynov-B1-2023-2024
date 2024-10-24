package piscine

// AppendRange prend deux entiers min et max et retourne une tranche d'entiers
// contenant tous les nombres de min (inclus) à max (exclus).
func AppendRange(min, max int) []int {
	var ans []int
	// Boucle de min à max (exclus) et ajoute chaque entier à la tranche ans.
	for i := min; i < max; i++ {
		ans = append(ans, i)
	}
	// Retourne la tranche d'entiers.
	return ans
}
