package main

import (
	"fmt"
)

func Prairies(inv *Inventory) {
	for {
		fmt.Println("==== Vous êtes dans les prairies. ====")
		fmt.Printf("")
		fmt.Println("[1] Récolter des matériaux")
		fmt.Println("[2] Explorer")
		fmt.Println("[3] Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			for i := 0; i < 3; i++ {
				if !inv.AddItem(Wood) {
					fmt.Println("Votre inventaire est plein.")
					break
				}
			}

			fmt.Println("Vous avez récolté 3 bois.")
		}

		if choix == 2 {
			fmt.Println("Vous explorez les prairies.")
		}

		if choix == 3 {
			return
		}
	}
}
