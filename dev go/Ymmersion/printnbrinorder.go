package piscine

import "github.com/01-edu/z01"

// PrintNbrInOrder affiche les chiffres d'un nombre dans l'ordre croissant
func PrintNbrInOrder(n int) {
	// Si le nombre est 0, on affiche '0' et on retourne
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var array []int
		eachValue := 0
		arrayCount := 0
		var minValue int

		// On extrait chaque chiffre du nombre et on les ajoute dans un tableau
		for n != 0 {
			eachValue = n % 10
			n /= 10
			array = append(array, eachValue)
		}

		// On compte le nombre d'éléments dans le tableau
		for count := range array {
			arrayCount = count + 1
		}

		// On trie le tableau en utilisant l'algorithme de tri à bulles
		for i := 0; i < arrayCount-1; i++ {
			for j := 0; j < arrayCount-i-1; j++ {
				if array[j] > array[j+1] {
					minValue = array[j]
					array[j] = array[j+1]
					array[j+1] = minValue
				}
			}
		}

		// On affiche chaque chiffre trié
		for i := 0; i < arrayCount; i++ {
			z01.PrintRune(rune(array[i] + 48))
		}
	}
}
