package main

import "fmt"

func Foret() {
	for {
		fmt.Println("FORET DE MARCHANG")
		fmt.Println("1. Récolter du bois")
		fmt.Println("2. Explorer")
		fmt.Println("3. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			bois = bois + 3
			fmt.Println("Vous avez récolté 3 bois.")
			fmt.Println("Vous avez", bois, "bois.")
		}

		if choix == 2 {
			fmt.Println("Vous explorez la forêt.")
		}

		if choix == 3 {
			return
		}
	}
}
