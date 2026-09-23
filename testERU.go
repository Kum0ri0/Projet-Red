package main

import "fmt"



type Sort struct {
	nom    string
	degats int
}

func sortsDuJoueur() []Sort {
	return []Sort{
		{nom: "Coup de poing", degats: 8},
		{nom: "Boule de Feu", degats: 18},
	}
}








func main() {

	
	perso := CreatePlayer("Subaru", "Humain")

	fmt.Println("=== AVANT ===")
	fmt.Println("Or :", perso.Gold, "| XP :", perso.XP, "| Niveau :", perso.Level)
	fmt.Println("Objets :", perso.Inventory.Size())
	fmt.Println()

	chooseZone(&perso)

	fmt.Println()
	fmt.Println("=== APRÈS ===")
	fmt.Println("Or :", perso.Gold, "| XP :", perso.XP, "| Niveau :", perso.Level)
	fmt.Println("Inventaire :")
	for _, it := range perso.Inventory.Items {
		fmt.Println(" -", it.Name)
	}
}