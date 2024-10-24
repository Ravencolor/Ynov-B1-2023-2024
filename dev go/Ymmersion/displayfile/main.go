package main

import (
	"fmt"
	"os"
)

func main() {
	// Vérifie si un argument est passé
	if len(os.Args) > 1 {
		// Ignore le premier argument (le nom du programme)
		os.Args = os.Args[1:]
	} else {
		// Affiche un message d'erreur si aucun fichier n'est spécifié
		fmt.Println("Nom de fichier manquant")
		return
	}

	// Vérifie s'il y a exactement un argument (le nom du fichier)
	if len(os.Args) == 1 {
		// Tente d'ouvrir le fichier spécifié
		file, err := os.Open(os.Args[0])
		if err != nil {
			// Affiche une erreur si le fichier ne peut pas être ouvert
			fmt.Println(err)
		} else {
			// Lit les 14 premiers octets du fichier
			data := make([]byte, 14)
			file.Read(data)
			// Affiche le contenu lu
			fmt.Println(string(data))
			// Ferme le fichier
			file.Close()
		}
	} else {
		// Affiche un message d'erreur si trop d'arguments sont passés
		fmt.Println("Trop d'arguments")
	}
}
