package main

import (
	"fmt"
)

func Medecin(joueur *Player) {
	for {
		fmt.Println("MEDECIN")
		fmt.Println("1. Se soigner")
		fmt.Println("2. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			if joueur.HP < joueur.MaxHP {
				joueur.Heal(joueur.MaxHP)
				fmt.Println("Vous êtes maintenant à", joueur.HP, "PV.")
			} else {
				fmt.Println("Vous avez déjà tous vos PV.")
			}
		}

		if choix == 2 {
			return
		}
	}
}
