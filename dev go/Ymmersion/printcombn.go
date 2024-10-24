package piscine

import "github.com/01-edu/z01"

// Convertit un entier en chaîne de caractères (0-9)
func IntToString(num int) string {
	if num == 0 {
		return "0"
	} else if num == 1 {
		return "1"
	} else if num == 2 {
		return "2"
	} else if num == 3 {
		return "3"
	} else if num == 4 {
		return "4"
	} else if num == 5 {
		return "5"
	} else if num == 6 {
		return "6"
	} else if num == 7 {
		return "7"
	} else if num == 8 {
		return "8"
	} else {
		return "9"
	}
}

// Convertit un entier en chaîne de caractères
func itoa(num int) string {
	var result string
	if num == 0 {
		return "0"
	}
	for num > 0 {
		result = IntToString(num%10) + result
		num /= 10
	}
	return result
}

// Affiche une chaîne de caractères
func printString(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

// Fonction récursive pour générer et afficher les combinaisons
func printComb(n int, prev int, result string, count *int) {
	for i := 0; i < 10; i++ {
		if prev < i {
			if n == 1 {
				if *count > 0 {
					printString(", ")
				}
				printString(result + itoa(i))
				*count++
			} else {
				printComb(n-1, i, result+itoa(i), count)
			}
		}
	}
}

// Fonction principale pour démarrer la génération des combinaisons
func PrintCombN(n int) {
	var count int = 0
	for i := 0; i < 10; i++ {
		if n > 1 {
			printComb(n-1, i, itoa(i), &count)
		} else {
			if count > 0 {
				printString(", ")
			}
			printString(itoa(i))
			count++
		}
	}
	printString("\n")
}
