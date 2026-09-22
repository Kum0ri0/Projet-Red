package main

import "fmt"

func Forgeron() {
	fmt.Println("Vous êtes chez le forgeron.")
	fmt.Println("1. Fabriquer une épée")
    fmt.Println("2. Retourner au village")

	var choix int
    fmt.Scan(&choix)

	if choix == 1 {
	if bois >= 3 {
		fmt.Println("Vous avez assez de bois.")
	} else {
		fmt.Println("Vous n'avez pas assez de bois.")
	}
}
}
