package piscine

import "github.com/01-edu/z01"

// PrintNum affiche un nombre entier en utilisant des runes
func PrintNum(num int) {
	i := '0'
	// Si le nombre est 0, on affiche '0' et on retourne
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	// Boucle pour ajuster la rune 'i' en fonction du dernier chiffre du nombre
	for j := 1; j <= num%10; j++ {
		i++
	}
	for j := -1; j >= num%10; j-- {
		i++
	}
	// Appel récursif pour afficher les chiffres restants
	if num/10 != 0 {
		PrintNum(num / 10)
	}
	// Affiche la rune correspondante au chiffre
	z01.PrintRune(i)
	return
}

// PrintNbr affiche un nombre entier, avec un signe négatif si nécessaire
func PrintNbr(n int) {
	// Si le nombre est négatif, on affiche '-'
	if n < 0 {
		z01.PrintRune('-')
	}
	// Appel de PrintNum pour afficher le nombre
	PrintNum(n)
}
