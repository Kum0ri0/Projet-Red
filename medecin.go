package main

import "fmt"

func Medecin() {
	for {
		fmt.Println("MEDECIN")
		fmt.Println("1. Se soigner")
		fmt.Println("2. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			if vie < vieMax {
				vie = vieMax
				fmt.Println("Vous êtes maintenant à", vie, "PV.")
			} else {
				fmt.Println("Vous avez déjà tous vos PV.")
			}
		}

		if choix == 2 {
			return
		}
	}
}