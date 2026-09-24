
package main

import (
	"fmt"
)

const MaxInventorySize = 10

type Inventory struct {
	Items []Item
}

// Ajouter un objet dans l'inventaire
func (i *Inventory) AddItem(newItem Item) bool {
	if len(i.Items) >= MaxInventorySize {
		return false
	}

	i.Items = append(i.Items, newItem)
	return true
}

// Supprimer un objet avec son index
func (i *Inventory) RemoveItem(index int) bool {
	if index < 0 || index >= len(i.Items) {
		return false
	}

	i.Items = append(i.Items[:index], i.Items[index+1:]...)
	return true
}

// Vérifier si l'inventaire possède un objet
func (i *Inventory) HasItem(target Item) bool {
	for _, currentItem := range i.Items {
		if currentItem.Name == target.Name {
			return true
		}
	}

	return false
}

// Vérifier si l'inventaire est plein
func (i Inventory) IsFull() bool {
	return len(i.Items) >= MaxInventorySize
}

// Retourner le nombre d'objets
func (i Inventory) Size() int {
	return len(i.Items)
}

// Afficher l'inventaire
func (i Inventory) Display() {
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║           INVENTAIRE            ║")
	fmt.Println("╠════════════════════════════════╣")

	if len(i.Items) == 0 {
		fmt.Println("║ Inventaire vide                ║")
	} else {
		for index, currentItem := range i.Items {
			fmt.Printf("║ [%d] %-27s ║\n", index, currentItem.Name)
		}
	}

	fmt.Println("╠════════════════════════════════╣")
	fmt.Printf("║ Objets : %d / %d                 ║\n", len(i.Items), MaxInventorySize)
	fmt.Println("╚════════════════════════════════╝")
}

// Utiliser un objet
func (i *Inventory) UseItem(index int, joueur *Player) bool {
	if index < 0 || index >= len(i.Items) {
		fmt.Println("Objet invalide.")
		return false
	}

	objet := i.Items[index]

	switch objet.Name {

	case "Potion de soin":
		if joueur.HP >= joueur.MaxHP {
			fmt.Println("Vous avez déjà tous vos PV.")
			return false
		}

		joueur.Heal(50)

		fmt.Println("Vous avez utilisé une potion de soin !")
		fmt.Println("PV :", joueur.HP, "/", joueur.MaxHP)

	case "Potion de poison":
		joueur.TakeDamage(30)

		fmt.Println("Vous avez utilisé une potion de poison !")
		fmt.Println("PV :", joueur.HP, "/", joueur.MaxHP)

	default:
		fmt.Println("Cet objet ne peut pas être utilisé.")
		return false
	}

	// Retirer l'objet après utilisation
	i.RemoveItem(index)

	return true
}
