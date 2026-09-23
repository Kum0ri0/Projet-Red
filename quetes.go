package main

import (
	"fmt"
)

var boisQuete int
var queteBoisTerminee bool

var metalQuete int
var queteMetalTerminee bool

var epeeBoisQuete int
var epeeMetalQuete int

var queteEpeeBoisTerminee bool
var queteEpeeMetalTerminee bool

func Quetes(joueur *Player) {
	for {
		fmt.Println("QUÊTES")
		fmt.Println("1. Le bois du village")
		fmt.Println("2. Le métal de la mine")
		fmt.Println("3. Fabriquer une épée en bois")
		fmt.Println("4. Fabriquer une épée en métal")
		fmt.Println("5. Retourner au village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			if queteBoisTerminee {
				fmt.Println("Cette quête est déjà terminée.")
			} else {
				fmt.Println("QUÊTE : Le bois du village")
				fmt.Println("Objectif : récolter 10 bois.")
				fmt.Println("Progression :", boisQuete, "/ 10")

				if boisQuete >= 10 {
					queteBoisTerminee = true
					joueur.AddGold(50)

					fmt.Println("QUÊTE TERMINÉE !")
					fmt.Println("Vous gagnez 50 pièces.")
				}
			}
		}

		if choix == 2 {
			if queteMetalTerminee {
				fmt.Println("Cette quête est déjà terminée.")
			} else {
				fmt.Println("QUÊTE : Le métal de la mine")
				fmt.Println("Objectif : récolter 10 métaux.")
				fmt.Println("Progression :", metalQuete, "/ 10")

				if metalQuete >= 10 {
					queteMetalTerminee = true
					joueur.AddGold(75)

					fmt.Println("QUÊTE TERMINÉE !")
					fmt.Println("Vous gagnez 75 pièces.")
				}
			}
		}

		if choix == 3 {
			if queteEpeeBoisTerminee {
				fmt.Println("Cette quête est déjà terminée.")
			} else {
				fmt.Println("QUÊTE : L'épée en bois")
				fmt.Println("Objectif : fabriquer 1 épée en bois.")
				fmt.Println("Progression :", epeeBoisQuete, "/ 1")

				if epeeBoisQuete >= 1 {
					queteEpeeBoisTerminee = true
					joueur.AddGold(50)

					fmt.Println("QUÊTE TERMINÉE !")
					fmt.Println("Vous gagnez 50 pièces.")
				}
			}
		}

		if choix == 4 {
			if queteEpeeMetalTerminee {
				fmt.Println("Cette quête est déjà terminée.")
			} else {
				fmt.Println("QUÊTE : L'épée en métal")
				fmt.Println("Objectif : fabriquer 1 épée en métal.")
				fmt.Println("Progression :", epeeMetalQuete, "/ 1")

				if epeeMetalQuete >= 1 {
					queteEpeeMetalTerminee = true
					joueur.AddGold(100)

					fmt.Println("QUÊTE TERMINÉE !")
					fmt.Println("Vous gagnez 100 pièces.")
				}
			}
		}

		if choix == 5 {
			return
		}
	}
}
