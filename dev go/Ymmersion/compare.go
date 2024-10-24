package piscine

// Compare compare deux chaînes de caractères a et b.
// Retourne 0 si a est égal à b, -1 si a est inférieur à b, et 1 si a est supérieur à b.
func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}
