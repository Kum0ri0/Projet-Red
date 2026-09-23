package main

import "fmt"

func main() {
	rudeus := CreatePlayer("Rudeus", "Mage")

	rudeus.Display()

	rudeus.SpellBook()

	DisplayAll()

	inv := Inventory{}

	inv.AddItem(HealingPotion)
	inv.AddItem(Sword)
	inv.AddItem(Fur)

	inv.Display()

	var choix int

	fmt.Print("Choisissez un objet : ")
	fmt.Scan(&choix)

	if inv.RemoveItem(choix) {
		fmt.Println("Objet retiré de l'inventaire.")
	} else {
		fmt.Println("Choix invalide.")
	}

	inv.Display()
}

