package piscine

// Any applique la fonction f à chaque élément de arr.
// Si f retourne true pour au moins un élément, Any retourne true.
// Sinon, Any retourne false.
func Any(f func(string) bool, arr []string) bool {
	for _, res := range arr {
		if f(res) {
			return true
		}
	}
	return false
}
