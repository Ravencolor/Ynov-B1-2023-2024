package main

import "github.com/01-edu/z01"

// Définition de la structure point avec deux champs x et y
type point struct {
	x int
	y int
}

// Fonction pour définir les valeurs de x et y dans une structure point
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	// Chaînes de caractères pour l'affichage
	xStr := "x = "
	yStr := "y = "
	// Création d'un pointeur vers une structure point
	points := &point{}

	// Appel de la fonction pour définir les valeurs de x et y
	setPoint(points)

	// Tableaux de runes pour stocker les chiffres de x et y
	xmassiv := []rune{}
	ymassiv := []rune{}

	// Récupération des valeurs de x et y
	xVal := points.x
	yVal := points.y

	// Conversion de la valeur de x en runes
	for xVal != 0 {
		xmassiv = append(xmassiv, rune(xVal%10))
		xVal /= 10
	}
	// Affichage de la chaîne "x = "
	for _, val := range xStr {
		z01.PrintRune(rune(val))
	}
	// Affichage des chiffres de x
	for i := len(xmassiv) - 1; i >= 0; i-- {
		z01.PrintRune(xmassiv[i] + 48)
	}
	// Affichage de la virgule et de l'espace
	z01.PrintRune(',')
	z01.PrintRune(' ')

	// Conversion de la valeur de y en runes
	for yVal != 0 {
		ymassiv = append(ymassiv, rune(yVal%10))
		yVal /= 10
	}
	// Affichage de la chaîne "y = "
	for _, val := range yStr {
		z01.PrintRune(rune(val))
	}
	// Affichage des chiffres de y
	for i := len(ymassiv) - 1; i >= 0; i-- {
		z01.PrintRune(ymassiv[i] + 48)
	}
	// Affichage du saut de ligne
	z01.PrintRune('\n')
}
