package piscine

// ConvertBase convertit une chaîne de caractères s d'une base t à une base p
func ConvertBase(s, t, p string) string {
	ln := 0                // Longueur de la chaîne s
	ln2 := 0               // Longueur de la base t
	ln3 := 0               // Longueur de la base p
	ans := ""              // Résultat final de la conversion
	st_t := map[rune]int{} // Map pour stocker les valeurs des caractères de la base t

	// Calcul de la longueur de la chaîne s
	for c := range s {
		ln = c + 1
	}

	// Initialisation de la map st_t avec les valeurs des caractères de la base t
	// et calcul de la longueur de la base t
	for i, c := range t {
		st_t[c] = i
		ln2 = i + 1
	}

	// Calcul de la longueur de la base p
	for c := range p {
		ln3 = c + 1
	}

	pw := 1  // Puissance de la base t
	cnt := 0 // Valeur numérique de la chaîne s dans la base t

	// Conversion de la chaîne s en une valeur numérique dans la base t
	for i := ln - 1; i >= 0; i-- {
		cnt = cnt + st_t[[]rune(s)[i]]*pw
		pw *= ln2
	}

	// Conversion de la valeur numérique cnt en une chaîne de caractères dans la base p
	for cnt != 0 {
		ans = string(p[cnt%ln3]) + ans
		cnt /= ln3
	}

	return ans
}
