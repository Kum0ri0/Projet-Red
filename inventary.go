package main

import (
	"fmt"
)

const MaxInventorySize = 10

type Inventory struct {
	Items []Item
}

// AddItem ajoute un objet à l'inventaire.
// Retourne false si l'inventaire est plein.
func (i *Inventory) AddItem(newItem Item) bool {
	if len(i.Items) >= MaxInventorySize {
		return false
	}

	i.Items = append(i.Items, newItem)
	return true
}

// RemoveItem retire un objet de l'inventaire.
// Retourne false si l'emplacement n'existe pas.
func (i *Inventory) RemoveItem(index int) bool {
	if index < 0 || index >= len(i.Items) {
		return false
	}

	i.Items = append(i.Items[:index], i.Items[index+1:]...)
	return true
}

// HasItem vérifie si un objet est présent dans l'inventaire.
func (i *Inventory) HasItem(target Item) bool {
	for _, currentItem := range i.Items {
		if currentItem.Name == target.Name {
			return true
		}
	}

	return false
}

// IsFull vérifie si l'inventaire est plein.
func (i Inventory) IsFull() bool {
	return len(i.Items) >= MaxInventorySize
}

// Size retourne le nombre d'objets présents.
func (i Inventory) Size() int {
	return len(i.Items)
}

// Display affiche le contenu de l'inventaire.
func (i Inventory) Display() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║          INVENTAIRE          ║")
	fmt.Println("╠══════════════════════════════╣")

	if len(i.Items) == 0 {
		fmt.Println("║ Inventaire vide              ║")
	} else {
		for index, currentItem := range i.Items {
			fmt.Printf("║ [%d] %-23s ║\n", index, currentItem.Name)
		}
	}

	fmt.Println("╠══════════════════════════════╣")
	fmt.Printf("║ Objets : %d / %d              ║\n", len(i.Items), MaxInventorySize)
	fmt.Println("╚══════════════════════════════╝")
}