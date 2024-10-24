package piscine

// IsPrime vérifie si un nombre est premier
func IsPrime(nb int) bool {
	// Si le nombre est positif
	if nb > 0 {
		// 1 n'est pas un nombre premier
		if nb == 1 {
			return false
		}
		// 2 et 3 sont des nombres premiers
		if nb == 2 || nb == 3 {
			return true
		}
		// Si le nombre est divisible par 2 ou 3, il n'est pas premier
		if nb%3 == 0 || nb%2 == 0 {
			return false
		}
		// Vérifie les autres diviseurs possibles
		for a := 2; a < nb; a++ {
			if nb%a == 0 {
				return false
			}
		}
		// Si aucun diviseur n'a été trouvé, le nombre est premier
		return true
	} else {
		// Si le nombre est négatif ou zéro, il n'est pas premier
		return false
	}
}
