package main

import "fmt"

func main() {
	QuadD(5, 3)
}

// QuadD dessine un rectangle de dimensions x et y avec des caractères spécifiques
func QuadD(x, y int) {
	if x > 0 && y > 0 {
		// Affiche la première ligne du rectangle
		fmt.Println("A" + repeat("B", x-2) + "C")
		// Affiche les lignes intermédiaires du rectangle
		for i := 1; i <= y-2; i++ {
			fmt.Print("B")
			for j := 1; j <= x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("B")
			fmt.Println()
		}
		// Affiche la dernière ligne du rectangle
		fmt.Println("A" + repeat("B", x-2) + "C")
		fmt.Println()
	} else {
		// Affiche un message d'erreur si les dimensions ne sont pas positives
		fmt.Println("Les dimensions doivent être positives.")
		return
	}
}

// repeat répète la chaîne s, n fois
func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
