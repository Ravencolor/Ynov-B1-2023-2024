package piscine

// IterativePower calcule nb élevé à la puissance power de manière itérative
func IterativePower(nb int, power int) int {
	a := nb
	// Si la puissance est négative, retourne 0
	if power < 0 {
		return 0
	} else {
		// Si la puissance est 0, retourne 1 (nb^0 = 1)
		if power == 0 {
			return 1
		}
		// Boucle pour multiplier nb par lui-même power fois
		for i := power; i > 1; i-- {
			a = nb * a
		}
		return a
	}
}
