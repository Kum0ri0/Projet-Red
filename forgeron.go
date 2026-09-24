package main

import "fmt"

func Forgeron(inv *Inventory) {
	for {
		fmt.Println("===== FORGERON =====")
		fmt.Println()
		fmt.Println("[1] Fabriquer une épée en bois")
		fmt.Println("[2] Fabriquer une épée en métal")
		fmt.Println("[3] Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			bois := compterItem(inv, "Bois")

			if bois >= 5 {
				for i := 0; i < 5; i++ {
					supprimerItem(inv, "Bois")
				}

				if inv.AddItem(WoodenSword) {
					fmt.Println("Vous avez fabriqué une épée en bois.")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de bois.")
			}
		}

		if choix == 2 {
			metal := compterItem(inv, "Métal")
			bois := compterItem(inv, "Bois")

			if metal >= 3 && bois >= 1 {
				for i := 0; i < 3; i++ {
					supprimerItem(inv, "Métal")
				}

				supprimerItem(inv, "Bois")

				if inv.AddItem(MetalSword) {
					fmt.Println("Vous avez fabriqué une épée en métal.")
				}
			} else {
				fmt.Println("Vous n'avez pas assez de matériaux.")
			}
		}

		if choix == 3 {
			return
		}
	}
}

func compterItem(inv *Inventory, nom string) int {
	compteur := 0

	for _, objet := range inv.Items {
		if objet.Name == nom {
			compteur++
		}
	}

	return compteur
}

func supprimerItem(inv *Inventory, nom string) bool {
	for index, objet := range inv.Items {
		if objet.Name == nom {
			return inv.RemoveItem(index)
		}
	}

	return false
}
