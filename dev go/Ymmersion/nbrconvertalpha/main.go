package main

import (
	"os"

	"github.com/01-edu/z01"
)

// PerfAtoi convertit une chaîne de caractères représentant un nombre en entier.
// Retourne 0 si la chaîne contient des caractères non numériques.
func PerfAtoi(s string) int {
	var i int = 0
	for _, c := range s {
		if !(c >= '0' && c <= '9') {
			return 0
		}
		i = i*10 + int(c-'0')
	}
	return i
}

// Atoi convertit une chaîne de caractères en entier, en gérant les signes '+' et '-'.
// Retourne 0 si la chaîne est vide ou contient des caractères non valides.
func Atoi(s string) int {
	var i int = 0
	if len(s) == 0 || s == "-" || s == "+" {
		return i
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || c == '-' || c == '+') {
			return i
		}
	}
	if s[0] == '-' && s[1] == '-' {
		return i
	}
	if s[0] == '-' && len(s) != 1 {
		s = s[1:]
		i = PerfAtoi(s)
		i = -i
	} else if s[0] == '+' && len(s) != 1 {
		s = s[1:]
		i = PerfAtoi(s)
	} else {
		i = PerfAtoi(s)
	}
	return i
}

func main() {
	args := os.Args[1:]
	caps := false
	if len(args) == 0 {
		return
	}
	// Vérifie si le premier argument est "--upper" pour activer les majuscules.
	if args[0] == "--upper" {
		caps = true
	}
	for i, arg := range args {
		if i == 0 && caps {
			continue
		}
		char_int := Atoi(arg)
		// Si la conversion échoue ou si le nombre est hors de l'intervalle [1, 26], imprime un espace.
		if char_int == 0 || char_int > 26 {
			z01.PrintRune(' ')
		} else {
			// Imprime la lettre correspondante en majuscule ou minuscule selon le flag caps.
			if caps {
				z01.PrintRune(rune(char_int + 'A' - 1))
			} else {
				z01.PrintRune(rune(char_int + 'a' - 1))
			}
		}
	}
	z01.PrintRune('\n')
}
