package main 

import "fmt"

func Prairies() {
	for {
		fmt.Println("Vous êtes dans les prairies.")
		fmt.Println("1. Récolter des matériaux")
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
			fmt.Println("Vous explorez les prairies.")
		}

		if choix == 3 {
			return
		}
	}
}