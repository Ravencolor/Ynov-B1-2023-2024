package piscine

// DivMod calcule la division et le reste de la division de a par b.
// Les résultats sont stockés dans les pointeurs div et mod.
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b // Stocke le résultat de la division entière de a par b dans div.
	*mod = a % b // Stocke le reste de la division de a par b dans mod.
}
