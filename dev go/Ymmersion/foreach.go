package piscine

// ForEach applique la fonction f à chaque élément du tableau arr.
func ForEach(f func(int), arr []int) {
	for _, res := range arr {
		f(res)
	}
}
