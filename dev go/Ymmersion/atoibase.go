// AtoiBase convertit une chaîne de caractères `s` en un entier en utilisant une base définie par la chaîne `t`.
// La fonction retourne 0 si `t` contient des caractères dupliqués ou les caractères '-' ou '+'.
//
// Paramètres:
// - s: La chaîne de caractères à convertir.
// - t: La chaîne de caractères représentant la base.
//
// Retourne:
// - Un entier représentant la valeur de `s` dans la base définie par `t`.
//
// Exemple:
// - AtoiBase("101", "01") retourne 5 car "101" en base binaire (base 2) est égal à 5 en base décimale.
//
// Note:
// - La fonction ne gère pas les bases avec un seul caractère ou les bases contenant des caractères invalides.
package piscine

func AtoiBase(s string, t string) int {
	res := 0
	lens := 0
	arr := map[rune]bool{}
	for _, c := range t {
		if arr[c] || c == '-' || c == '+' {
			return res
		}
		arr[c] = true
		lens++
	}
	if lens > 1 {
		for _, v := range s {
			count := 0
			if arr[v] {
				for _, j := range t {
					if j == v {
						break
					}
					count++
				}
				res = res*lens + count
			}
		}
	}
	return res
}
