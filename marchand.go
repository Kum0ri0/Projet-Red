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

		if choix == 2 {
			fmt.Println("VENDRE")
			fmt.Println("1. Vendre du bois - 5 pièces")
			fmt.Println("2. Vendre du métal - 10 pièces")
			fmt.Println("3. Vendre une épée en bois - 20 pièces")
			fmt.Println("4. Vendre une épée en métal - 40 pièces")
			fmt.Println("5. Retour")

			var vente int
			fmt.Scan(&vente)

			if vente == 1 {
				if bois >= 1 {
					bois = bois - 1
					argent = argent + 5
					fmt.Println("Vous avez vendu 1 bois.")
				} else {
					fmt.Println("Vous n'avez pas de bois.")
				}
			}

			if vente == 2 {
				if metal >= 1 {
					metal = metal - 1
					argent = argent + 10
					fmt.Println("Vous avez vendu 1 métal.")
				} else {
					fmt.Println("Vous n'avez pas de métal.")
				}
			}

			if vente == 3 {
				if epeeBois >= 1 {
					epeeBois = epeeBois - 1
					argent = argent + 20
					fmt.Println("Vous avez vendu une épée en bois.")
				} else {
					fmt.Println("Vous n'avez pas d'épée en bois.")
				}
			}

			if vente == 4 {
				if epeeMetal >= 1 {
					epeeMetal = epeeMetal - 1
					argent = argent + 40
					fmt.Println("Vous avez vendu une épée en métal.")
				} else {
					fmt.Println("Vous n'avez pas d'épée en métal.")
				}
			}
		}

		if choix == 3 {
			return
		}
	}
}