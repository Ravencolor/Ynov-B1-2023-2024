package main

import "fmt"

func main() {
	QuadA(5, 3)
}

// QuadA dessine un rectangle de largeur x et de hauteur y
func QuadA(x, y int) {
	if x > 0 && y > 0 {
		// Affiche la première ligne du rectangle
		fmt.Println("o" + repeat("-", x-2) + "o")
		// Affiche les lignes intermédiaires du rectangle
		for i := 1; i <= y-2; i++ {
			fmt.Print("|")
			for j := 1; j <= x-2; j++ {
				fmt.Print(" ")
			}
			fmt.Print("|")
			fmt.Println()
		}
		// Affiche la dernière ligne du rectangle
		fmt.Println("o" + repeat("-", x-2) + "o")
		fmt.Println()
	} else {
		// Affiche un message d'erreur si les dimensions sont négatives ou nulles
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
