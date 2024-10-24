package main

import "fmt"

func main() {
	QuadC(5, 3)
}

// QuadC dessine un rectangle avec des dimensions x et y
func QuadC(x, y int) {
	if x > 0 && y > 0 {
		// Affiche la première ligne avec des 'A' aux extrémités et des 'B' au milieu
		fmt.Println("A" + repeat("B", x-2) + "A")
		// Affiche les lignes intermédiaires avec des 'B' aux extrémités et des espaces au milieu
		for i := 1; i <= y-2; i++ {
			fmt.Print("B")
			for j := 1; j <= x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("B")
			fmt.Println()
		}
		// Affiche la dernière ligne avec des 'C' aux extrémités et des 'B' au milieu
		fmt.Println("C" + repeat("B", x-2) + "C")
		fmt.Println()
	} else {
		// Affiche un message d'erreur si les dimensions ne sont pas positives
		fmt.Println("Les dimensions doivent être positives.")
		return
	}
}

// repeat répète une chaîne de caractères s, n fois
func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
