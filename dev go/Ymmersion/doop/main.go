package main

import (
	// "fmt"
	"os"
)

// Fonction pour imprimer une chaîne de caractères sur la console
func PrintConsole(str string) {
	os.Stdout.WriteString(str)
	os.Stdout.WriteString("\n")
	os.Stdout.Close()
}

// Fonction récursive pour convertir un nombre en chaîne de caractères
func NbrToStrRec(n, dot int64) string {
	if 10 > n && n > -1*10 {
		return string('0' + n*dot)
	}
	return NbrToStrRec(n/10, dot) + string('0'+(n%10)*dot)
}

// Fonction pour convertir un nombre en chaîne de caractères
func NbrToStr(n int64) string {
	dot := int64(1)
	res := ""
	if n == 0 {
		return "0"
	}
	if n < 0 {
		dot *= -1
		res += "-"
	}
	return res + NbrToStrRec(n, dot)
}

// Fonction pour vérifier le débordement lors de la conversion d'une chaîne en nombre
func AtoiOverflow(a, b, c int64) bool {
	if a < 0 && c < 0 {
		return a*b+c < 0
	} else if a > 0 && c > 0 {
		return a*b+c > 0
	}
	return true
}

// Fonction pour vérifier le débordement lors de la multiplication
func MultiplyOverflow(a, b, c int64) bool {
	prod := a*b + c
	return (prod/b)-c == a
}

// Fonction pour vérifier le débordement lors de l'addition
func PlusOverflow(a, b int64) bool {
	if a < 0 && b < 0 {
		return a+b < 0
	} else if a > 0 && b > 0 {
		return a+b > 0
	}
	return true
}

// Fonction pour vérifier le débordement lors de la soustraction
func MinusOverflow(a, b int64) bool {
	if a < 0 && b < 0 {
		if b <= a {
			return a-b >= 0
		}
		return a-b < 0
	} else if a > 0 && b > 0 {
		if a <= b {
			return a-b <= 0
		}
		return a-b > 0
	}
	return true
}

// Fonction pour convertir une chaîne de caractères en nombre
func Atoi(nbr string) (int64, bool) {
	var res int64 = 0
	var sign int64 = 1
	if nbr[0] == '-' {
		nbr = nbr[1:]
		sign *= -1
	} else if nbr[0] == '+' {
		nbr = nbr[1:]
	}
	for _, digit := range nbr {
		tmp := int64(digit-'0') * sign
		if !AtoiOverflow(res, int64(10), tmp) {
			return 0, false
		}
		res = res*10 + tmp
	}
	return res, true
}

// Fonction pour additionner deux nombres représentés en chaînes de caractères
func Plus(a, b string) {
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if !PlusOverflow(aa, bb) {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa + bb))
}

// Fonction pour soustraire deux nombres représentés en chaînes de caractères
func Deduct(a, b string) {
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if !MinusOverflow(aa, bb) {
		PrintConsole("-0")
		return
	}
	PrintConsole(NbrToStr(aa - bb))
}

// Fonction pour diviser deux nombres représentés en chaînes de caractères
func Devide(a, b string) {
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if bb == 0 {
		PrintConsole("No division by 0")
		return
	}
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa / bb))
}

// Fonction pour multiplier deux nombres représentés en chaînes de caractères
func Multiply(a, b string) {
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if !MultiplyOverflow(aa, bb, 0) {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa * bb))
}

// Fonction pour calculer le modulo de deux nombres représentés en chaînes de caractères
func Mod(a, b string) {
	bb, bBool := Atoi(b)
	if !bBool {
		PrintConsole("0")
		return
	}
	if bb == 0 {
		PrintConsole("No modulo by 0")
		return
	}
	aa, aBool := Atoi(a)
	if !aBool {
		PrintConsole("0")
		return
	}
	PrintConsole(NbrToStr(aa % bb))
}

// Fonction pour vérifier si une chaîne de caractères est numérique
func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}
	for _, s := range str {
		if s < 48 || s > 57 {
			return false
		}
	}
	return true
}

// Fonction principale
func main() {
	args := os.Args[1:]
	argsCount := 0
	for range args {
		argsCount++
	}
	if argsCount != 3 {
		return
	}
	if !(IsNumeric(args[0]) && IsNumeric(args[2])) {
		PrintConsole("0")
	}
	funcsArr := []func(string, string){Plus, Deduct, Devide, Multiply, Mod}
	operators := []string{"+", "-", "/", "*", "%"}
	for i, val := range operators {
		if val == args[1] {
			funcsArr[i](args[0], args[2])
			return
		}
	}
	PrintConsole("0")
}
