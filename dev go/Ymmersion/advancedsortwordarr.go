package piscine

// AdvancedSortWordArr trie un tableau de chaînes de caractères en utilisant une fonction de comparaison personnalisée.
// a : tableau de chaînes de caractères à trier
// f : fonction de comparaison qui prend deux chaînes de caractères et retourne un entier
//     - retourne 1 si la première chaîne doit être après la deuxième
//     - retourne -1 si la première chaîne doit être avant la deuxième
//     - retourne 0 si les deux chaînes sont égales
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			// Si la fonction de comparaison retourne 1, on échange les éléments
			if f(a[i], a[j]) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
