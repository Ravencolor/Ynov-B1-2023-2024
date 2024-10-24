package piscine

// ConcatParams prend une slice de chaînes de caractères en entrée
// et les concatène en une seule chaîne, séparée par des sauts de ligne.
func ConcatParams(args []string) string {
	myStr := "" // Chaîne résultante
	ln := 0     // Longueur de la slice

	// Boucle pour déterminer la longueur de la slice
	for i := range args {
		ln = i
	}

	// Boucle pour concaténer les chaînes avec des sauts de ligne
	for i := 0; i <= ln; i++ {
		if i != ln {
			myStr = myStr + args[i] + "\n" // Ajoute un saut de ligne entre les éléments
		} else {
			myStr = myStr + args[i] // Pas de saut de ligne après le dernier élément
		}
	}
	return myStr // Retourne la chaîne concaténée
}
