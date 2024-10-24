package piscine

// UltimateDivMod prend deux pointeurs vers des entiers a et b.
// Elle calcule le quotient de a divisé par b et le reste de a divisé par b.
// Ensuite, elle met à jour la valeur pointée par a avec le quotient et la valeur pointée par b avec le reste.
func UltimateDivMod(a *int, b *int) {
	quotient := *a / *b
	remainder := *a % *b
	*a = quotient
	*b = remainder
}
