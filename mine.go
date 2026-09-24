package main

import "fmt"

func Mine(inv *Inventory) {
	for {
		fmt.Println("===== MINE DE MARCHANG =====")
		fmt.Printf("")
		fmt.Println("[1] Miner du métal")
		fmt.Println("[2] Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			recolte := 0
			for i := 0; i < 2; i++ {
				if !inv.AddItem(Metal) {
					fmt.Println("Votre inventaire est plein.")
					break
				}
				recolte++
				metalQuete++
			}

			fmt.Println("Vous avez récolté", recolte, "métaux.")
		}

		if choix == 2 {
			return
		}
	}
}
