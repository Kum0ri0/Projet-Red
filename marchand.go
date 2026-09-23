package main

import "fmt"

func Marchand() {
	for {
		fmt.Println("MARCHAND")
		fmt.Println("1. Acheter")
		fmt.Println("2. Vendre")
		fmt.Println("3. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			fmt.Println("ACHETER")
			fmt.Println("1. Acheter du bois - 10 pièces")
			fmt.Println("2. Acheter du métal - 20 pièces")
			fmt.Println("3. Retour")

			var achat int
			fmt.Scan(&achat)

			if achat == 1 {
				if argent >= 10 {
					argent = argent - 10
					bois = bois + 1
					fmt.Println("Vous avez acheté 1 bois.")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent.")
				}
			}

			if achat == 2 {
				if argent >= 20 {
					argent = argent - 20
					metal = metal + 1
					fmt.Println("Vous avez acheté 1 métal.")
				} else {
					fmt.Println("Vous n'avez pas assez d'argent.")
				}
			}
		}

		if choix == 3 {
			return
		}
	}
}