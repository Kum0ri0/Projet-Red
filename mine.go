package main

import "fmt"

func Mine(inv *Inventory) {
	for {
		fmt.Println("MINE DE MARCHANG")
		fmt.Println("1. Miner du métal")
		fmt.Println("2. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			for i := 0; i < 2; i++ {
				if !inv.AddItem(Metal) {
					fmt.Println("Votre inventaire est plein.")
					break
				}
			}

			fmt.Println("Vous avez récolté 2 métaux.")
		}

		if choix == 2 {
			return
		}
	}
}
