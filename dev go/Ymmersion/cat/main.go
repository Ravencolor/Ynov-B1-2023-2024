package main

import (
	"io/ioutil"
	"os"
)

// main est la fonction principale qui implémente une version simplifiée de la commande UNIX `cat`.
// Elle lit les fichiers passés en arguments et affiche leur contenu sur la sortie standard.
// Si aucun fichier n'est passé en argument, elle lit depuis l'entrée standard.
//
// Usage :
//
//	go run main.go [fichier1] [fichier2] ...
//
// Comportement :
// - Si aucun argument n'est fourni, la fonction lit depuis l'entrée standard (stdin) et affiche le contenu.
// - Si des arguments sont fournis, la fonction lit chaque fichier et affiche leur contenu.
// - En cas d'erreur lors de la lecture d'un fichier ou de l'entrée standard, un message d'erreur est affiché sur la sortie d'erreur standard (stderr) et le programme se termine avec un code de sortie 1.
func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		} else {
			os.Stdout.Write(data)
		}
	} else {
		for _, s := range args {
			file, err := os.Open(s)
			if err != nil {
				os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
				break
			} else {
				data, err := ioutil.ReadAll(file)
				if err != nil {
					os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
					os.Exit(1)
					break
				} else {
					os.Stdout.Write(data)
				}
			}
		}
	}
}
