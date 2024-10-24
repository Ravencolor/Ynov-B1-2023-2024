package piscine

// SplitWhiteSpaces sépare une chaîne de caractères en utilisant les espaces blancs comme délimiteurs
func SplitWhiteSpaces(str string) []string {
	a := 0
	ln := 0

	ok2 := false
	// Boucle pour compter le nombre de mots dans la chaîne
	for c := range str {
		// Si on trouve un caractère non blanc après un caractère blanc, on incrémente le compteur de mots
		if ok2 && c != 0 && (str[c-1] == '\n' || str[c-1] == '\t' || str[c-1] == ' ') && str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ln++
		}
		// Si le caractère actuel n'est pas un espace blanc, on active le flag ok2
		if str[c] != '\n' && str[c] != '\t' && str[c] != ' ' {
			ok2 = true
		}
	}
	ln++

	// Création du tableau de chaînes de caractères pour stocker les mots
	ans := make([]string, ln)
	index := 0
	myStr := ""
	// Boucle pour remplir le tableau avec les mots
	for i, w := range str {
		// Si on trouve un espace blanc, on ajoute le mot actuel au tableau
		if w == '\n' || w == ' ' || w == '\t' {
			if myStr != "" {
				ans[index] = myStr
				index++
				myStr = ""
				a = i
			}
		} else {
			// Ajout du caractère actuel au mot en cours
			if w != ' ' {
				myStr = myStr + string(w)
			}
		}
	}

	// Ajout du dernier mot au tableau si nécessaire
	if str[a+1:] != "" && str[a+1:] != " " {
		ans[index] = str[a+1:]
	}
	return ans
}
