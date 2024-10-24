package piscine

// RecursiveFactorial calcule la factorielle d'un nombre de manière récursive.
// Si le nombre est 0, la fonction retourne 1.
// Si le nombre est négatif ou trop grand (>= 1000000), la fonction retourne 0.
// Sinon, la fonction retourne nb multiplié par la factorielle de nb-1.
func RecursiveFactorial(nb int) int {
	a := 0
	if nb == 0 {
		return 1
	} else if nb < 0 || nb >= 1000000 {
		return 0
	} else {
		a = nb * RecursiveFactorial(nb-1)
	}
	return a
}
