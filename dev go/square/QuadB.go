package main

import "fmt"

func main() {
	QuadB(5, 3)
}

// QuadB dessine un rectangle de largeur x et de hauteur y
func QuadB(x, y int) {
	// Vérifie que les dimensions sont positives
	if x > 0 && y > 0 {
		// Affiche la première ligne du rectangle
		fmt.Println("/" + repeat("*", x-2) + "\\")
		// Affiche les lignes intermédiaires du rectangle
		for i := 1; i <= y-2; i++ {
			fmt.Print("*")
			for j := 1; j <= x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
			fmt.Println()
		}
		// Affiche la dernière ligne du rectangle
		fmt.Println("\\" + repeat("*", x-2) + "/")
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
