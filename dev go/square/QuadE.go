package main

import "fmt"

func main() {
	QuadE(5, 3)
}

// QuadE dessine un rectangle de dimensions x par y avec des caractères spécifiques.
func QuadE(x, y int) {
	if x > 0 && y > 0 {
		// Dessine la première ligne avec les coins "A" et "C" et les bords "B"
		fmt.Println("A" + repeat("B", x-2) + "C")
		// Dessine les lignes intermédiaires avec les bords "B" et des espaces au milieu
		for i := 1; i <= y-2; i++ {
			fmt.Print("B")
			for j := 1; j <= x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("B")
			fmt.Println()
		}
		// Dessine la dernière ligne avec les coins "C" et "A" et les bords "B"
		fmt.Println("C" + repeat("B", x-2) + "A")
		fmt.Println()
	} else {
		// Affiche un message d'erreur si les dimensions sont négatives ou nulles
		fmt.Println("Les dimensions doivent être positives.")
		return
	}
}

// repeat répète une chaîne de caractères s, n fois.
func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
