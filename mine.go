package main

import "fmt"

func Mine() {
	for {
		fmt.Println("MINE DE MARCHANG")
		fmt.Println("1. Miner du métal")
		fmt.Println("2. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			metal = metal + 2
			fmt.Println("Vous avez récolté 2 métaux.")
			fmt.Println("Vous avez", metal, "métaux.")
		}

		if choix == 2 {
			return
		}
	}
}
