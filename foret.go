package main

import "fmt"

func Foret(inv *Inventory) {
	for {
		fmt.Println("FORET DE MARCHANG")
		fmt.Println("1. Récolter du bois")
		fmt.Println("2. Explorer")
		fmt.Println("3. Retourner au village")

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
			fmt.Println("Vous explorez la forêt.")
		}

		if choix == 3 {
			return
		}
	}
}


