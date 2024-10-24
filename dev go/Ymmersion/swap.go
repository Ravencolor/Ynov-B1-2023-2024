package piscine

// Swap échange les valeurs de deux entiers pointés par a et b.
func Swap(a *int, b *int) {
	*a, *b = *b, *a
}
