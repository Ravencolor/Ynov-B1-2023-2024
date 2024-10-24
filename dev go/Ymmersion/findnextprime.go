package piscine

// FindNextPrime trouve le prochain nombre premier après ou égal à nb
func FindNextPrime(nb int) int {
	for a := nb; a >= nb; a++ {
		// Vérifie si le nombre actuel est premier
		if IsPrime(a) {
			// Retourne le nombre premier trouvé
			return a
		}
	}
	// Retourne 1 si aucun nombre premier n'est trouvé (ce cas ne devrait pas se produire)
	return 1
}
