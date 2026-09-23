package main

import "fmt"

func Forgeron() {
	for {
		fmt.Println("FORGERON")
		fmt.Println("1. Fabriquer une épée en bois")
		fmt.Println("2. Fabriquer une épée en métal")
		fmt.Println("3. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			if bois >= 5 {
				bois = bois - 5
				epeeBois = epeeBois + 1
				epeeBoisQuete = epeeBoisQuete + 1

				fmt.Println("Vous avez fabriqué une épée en bois.")
			} else {
				fmt.Println("Vous n'avez pas assez de bois.")
			}
		}

		if choix == 2 {
			if metal >= 3 && bois >= 1 {
				metal = metal - 3
				bois = bois - 1
				epeeMetal = epeeMetal + 1
				epeeMetalQuete = epeeMetalQuete + 1

				fmt.Println("Vous avez fabriqué une épée en métal.")
			} else {
				fmt.Println("Vous n'avez pas assez de matériaux.")
			}
		}

		if choix == 3 {
			return
		}
	}
}