package piscine

// Fibonacci calcule le n-ième nombre de Fibonacci.
// Si l'index est négatif, il retourne -1.
// Si l'index est 0, il retourne 0.
// Si l'index est 1, il retourne 1.
// Pour tout autre index, il retourne la somme des deux nombres de Fibonacci précédents.
func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}
