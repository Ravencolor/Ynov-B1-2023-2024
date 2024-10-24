package piscine

// StrLen retourne la longueur d'une chaîne de caractères en comptant les runes.
// Cela permet de gérer correctement les caractères spéciaux et les emojis.
func StrLen(ch string) int {
	return len([]rune(ch))
}
