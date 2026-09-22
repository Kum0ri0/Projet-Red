package main

import "fmt"

func Mine() {
	for {
		fmt.Println("Vous êtes dans la mine.")
		fmt.Println("1. Miner du métal")
		fmt.Println("2. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			fmt.Println("Vous minez du métal.")
		}

		if choix == 2 {
			return
		}
	}
}
