package piscine

// RecursivePower calcule nb élevé à la puissance power de manière récursive.
// Si power est 0, la fonction retourne 1.
// Si power est négatif, la fonction retourne 0.
// Sinon, la fonction retourne nb multiplié par le résultat de RecursivePower(nb, power-1).
func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
