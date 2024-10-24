package piscine

// IterativeFactorial calcule la factorielle d'un nombre de manière itérative.
// Si le nombre est négatif ou supérieur à 1 000 000, la fonction retourne 0.
// Si le nombre est 0, la fonction retourne 1.
func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 1000000 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return factorise(nb, 1) * nb
}

// factorise calcule la factorielle d'un nombre en utilisant une boucle for.
// La fonction prend deux paramètres : le nombre à factoriser et un index (non utilisé ici).
func factorise(n int, index int) int {
	fact := 1
	for i := 1; i < n; i++ {
		fact = fact * i
	}
	return fact
}
