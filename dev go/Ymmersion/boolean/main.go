package main

import (
	"os"

	"github.com/01-edu/z01"
)

// printStr affiche une chaîne de caractères, rune par rune
func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

// even vérifie si un nombre est pair
func even(nbr int) bool {
	return nbr%2 == 0
}

// isEven retourne vrai si le nombre est pair, sinon faux
func isEven(nbr int) bool {
	if even(nbr) {
		return true
	} else {
		return false
	}
}

func main() {
	// lengthOfArg contient le nombre d'arguments passés au programme
	lengthOfArg := len(os.Args[1:])
	if isEven(lengthOfArg) {
		EvenMsg := "I have an even number of arguments"
		printStr(EvenMsg)
	} else {
		OddMsg := "I have an odd number of arguments"
		printStr(OddMsg)
	}
}
