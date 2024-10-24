package piscine

// prim vérifie si le caractère est une lettre ou un chiffre
func prim(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

// Capitalize met en majuscule la première lettre de chaque mot et le reste en minuscule
func Capitalize(s string) string {
	ar := []rune(s)
	letra := true
	for i := 0; i < len(s); i++ {
		// Si le caractère est une lettre ou un chiffre et que c'est le début d'un mot
		if prim(ar[i]) == true && letra {
			// Si la lettre est en minuscule, la convertir en majuscule
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = 'A' - 'a' + ar[i]
			}
			letra = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			// Si la lettre est en majuscule, la convertir en minuscule
			ar[i] = 'a' - 'A' + ar[i]
		} else if prim(ar[i]) == false {
			// Si le caractère n'est pas une lettre ou un chiffre, le prochain caractère sera le début d'un mot
			letra = true
		}
	}
	return string(ar)
}
