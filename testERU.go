package main

import (
	"fmt"
	"time"
)

func main() {
	clearScreen()

	fmt.Println("=====================================")
	fmt.Println("        PROJET RED - COMBAT")
	fmt.Println("=====================================")
	fmt.Println()
	fmt.Print("Quel est ton nom ? ")

	var nom string
	fmt.Scanln(&nom)
	if nom == "" {
		nom = "Aventurier"
	}

	perso := CreatePlayer(nom, "Humain")

	// De quoi démarrer l'aventure.
	perso.Inventory.AddItem(HealingPotion)
	perso.Inventory.AddItem(HealingPotion)
	perso.Inventory.AddItem(PoisonPotion)
	perso.SpellBook()

	fmt.Println()
	fmt.Println("Bienvenue,", perso.Name, "!")
	fmt.Println("Tu pars avec 2 potions de soin, 1 potion de poison,")
	fmt.Println("et tu connais deja Coup de poing et Boule de Feu.")
	time.Sleep(3 * time.Second)

	for {
		clearScreen()

		perso.Display()

		fmt.Println()
		fmt.Println("=========== QUE FAIRE ? ===========")
		fmt.Println()
		fmt.Println("[1] Partir a l'aventure")
		fmt.Println("[2] Voir mon inventaire")
		fmt.Println("[3] Voir mes sorts")
		fmt.Println("[4] Se reposer (soin complet)")
		fmt.Println("[0] Quitter le jeu")
		fmt.Println()

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			clearScreen()
			chooseZone(&perso)

			if perso.HP <= 0 {
				fmt.Println()
				fmt.Println("Tu te reveilles au village, sauve de justesse...")
				perso.HP = perso.MaxHP / 2
			}

			fmt.Println()
			fmt.Println("Appuie sur Entree pour revenir au village.")
			fmt.Scanln()
			fmt.Scanln()

		case 2:
			clearScreen()
			perso.Inventory.Display()
			fmt.Println()
			fmt.Println("Appuie sur Entree pour continuer.")
			fmt.Scanln()
			fmt.Scanln()

		case 3:
			clearScreen()
			fmt.Println("=========== MES SORTS ===========")
			fmt.Println()
			for _, s := range perso.Spells {
				fmt.Println(" -", s.Name, ":", s.Damage, "degats")
			}
			fmt.Println()
			fmt.Println("Appuie sur Entree pour continuer.")
			fmt.Scanln()
			fmt.Scanln()

		case 4:
			perso.HP = perso.MaxHP
			fmt.Println()
			fmt.Println("Tu te reposes... PV restaures !")
			time.Sleep(2 * time.Second)

		case 0:
			clearScreen()
			fmt.Println("Merci d'avoir joue,", perso.Name, "!")
			fmt.Println("Niveau atteint :", perso.Level, "| Or :", perso.Gold)
			return

		default:
			fmt.Println("Ce choix n'existe pas...")
			time.Sleep(1 * time.Second)
		}
	}
}