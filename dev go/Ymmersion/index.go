package piscine

// Index retourne l'index de la première occurrence de t dans s, ou -1 si t n'est pas présent dans s.
func Index(s, t string) int {
	ln := 0
	ln2 := 0

	// Calcul de la longueur de la chaîne s
	for _, c := range s {
		if c == c {
			ln++
		}
	}

	// Calcul de la longueur de la chaîne t
	for _, c := range t {
		if c == c {
			ln2++
		}
	}

	// Parcours de la chaîne s pour trouver t
	for i := 0; i < ln; i++ {
		if ln2 != 0 && s[i] == t[0] {
			ok := true
			cur_ch := 0
			for j := 0; j < ln2; j++ {
				if i+cur_ch >= ln || t[j] != s[i+cur_ch] {
					ok = false
					break
				}
				cur_ch++
			}
			if ok == true {
				return i
			}
		}
	}

	// Si t est une chaîne vide, retourner 0
	if t == "" {
		return 0
	}

	// Si t n'est pas trouvé dans s, retourner -1
	return -1
}

// arrayCount retourne le nombre de caractères dans la chaîne s.
func arrayCount(s string) int {
	count := 0
	for i := range s {
		count = i + 1
	}
	return count
}
