package main

import (
	"fmt"
)

func Marchand(joueur *Player, inv *Inventory) {
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
				if joueur.Gold >= 10 {
					if inv.IsFull() {
						fmt.Println("Votre inventaire est plein.")
					} else {
						joueur.Gold -= 10
						inv.AddItem(Wood)
						fmt.Println("Vous avez acheté 1 bois.")
					}
				} else {
					fmt.Println("Vous n'avez pas assez d'argent.")
				}
			}

			if achat == 2 {
				if joueur.Gold >= 20 {
					if inv.IsFull() {
						fmt.Println("Votre inventaire est plein.")
					} else {
						joueur.Gold -= 20
						inv.AddItem(Metal)
						fmt.Println("Vous avez acheté 1 métal.")
					}
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
				index := trouverItem(inv, "Bois")

				if index != -1 {
					inv.RemoveItem(index)
					joueur.Gold += 5
					fmt.Println("Vous avez vendu 1 bois.")
				} else {
					fmt.Println("Vous n'avez pas de bois.")
				}
			}

			if vente == 2 {
				index := trouverItem(inv, "Métal")

				if index != -1 {
					inv.RemoveItem(index)
					joueur.Gold += 10
					fmt.Println("Vous avez vendu 1 métal.")
				} else {
					fmt.Println("Vous n'avez pas de métal.")
				}
			}

			if vente == 3 {
				index := trouverItem(inv, "Épée en bois")

				if index != -1 {
					inv.RemoveItem(index)
					joueur.Gold += 20
					fmt.Println("Vous avez vendu une épée en bois.")
				} else {
					fmt.Println("Vous n'avez pas d'épée en bois.")
				}
			}

			if vente == 4 {
				index := trouverItem(inv, "Épée en métal")

				if index != -1 {
					inv.RemoveItem(index)
					joueur.Gold += 40
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

func trouverItem(inv *Inventory, nom string) int {
	for index, objet := range inv.Items {
		if objet.Name == nom {
			return index
		}
	}

	return -1
}
