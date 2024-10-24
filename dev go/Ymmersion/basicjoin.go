package piscine

// BasicJoin prend une slice de chaînes de caractères et les concatène en une seule chaîne.
func BasicJoin(strs []string) string {
	var ret string
	// Parcourt chaque chaîne dans la slice et les ajoute à la chaîne de retour.
	for _, str := range strs {
		ret += str
	}
	// Retourne la chaîne concaténée.
	return ret
}
