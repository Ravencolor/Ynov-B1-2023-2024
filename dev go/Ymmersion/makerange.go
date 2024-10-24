package piscine

// MakeRange crée un tableau d'entiers contenant toutes les valeurs de min à max-1.
// Si min est supérieur ou égal à max, la fonction retourne nil.
func MakeRange(min, max int) []int {
	ln := max - min // Calculer la longueur du tableau
	var ans []int
	if min >= max {
		return nil // Retourner nil si min est supérieur ou égal à max
	}
	ans = make([]int, ln) // Initialiser le tableau avec la longueur calculée
	for i := 0; i < ln; i++ {
		ans[i] = min // Assigner la valeur min à la position i
		min++        // Incrémenter min
	}
	return ans // Retourner le tableau rempli
}
