package piscine

// PointOne prend un pointeur vers un entier en paramètre
// et modifie la valeur pointée pour qu'elle soit égale à 1.
func PointOne(n *int) {
	*n = 1
}
