package piscine

// UltimatePointOne met la valeur 1 à l'adresse pointée par le triple pointeur n
func UltimatePointOne(n ***int) {
	***n = 1
}
