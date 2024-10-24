package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		result = result*10 + int(s[i]-'0')
		i++
	}
	for i < len(s) && s[i] != ' ' {
		return 0
	}
	return result * sign
}

// Atoi convertit une chaîne de caractères en un entier.
// La fonction ignore les espaces blancs au début de la chaîne,
// gère les signes '+' et '-' et convertit les caractères numériques en entier.
// Si la chaîne contient des caractères non numériques après les chiffres,
// la fonction retourne 0.
