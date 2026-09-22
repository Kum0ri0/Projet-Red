package main

const MaxInventorySize = 10

type Inventory struct {
	Items []Item
}

// AddItem ajoute un objet à l'inventaire.
// Retourne false si l'inventaire est déjà plein.
func (i *Inventory) AddItem(item Item) bool {
	if len(i.Items) >= MaxInventorySize {
		return false
	}

	i.Items = append(i.Items, item)
	return true
}

// RemoveItem retire l'objet situé à l'emplacement indiqué.
// Retourne false si l'emplacement n'existe pas.
func (i *Inventory) RemoveItem(index int) bool {
	if index < 0 || index >= len(i.Items) {
		return false
	}

	i.Items = append(i.Items[:index], i.Items[index+1:]...)
	return true
}

// HasItem vérifie si un objet précis est présent.
func (i *Inventory) HasItem(item Item) bool {
	for _, currentItem := range i.Items {
		if currentItem.Name == item.Name {
			return true
		}
	}

	return false
}

// IsFull vérifie si l'inventaire est plein.
func (i *Inventory) IsFull() bool {
	return len(i.Items) >= MaxInventorySize
}

// Size retourne le nombre d'objets présents.
func (i *Inventory) Size() int {
	return len(i.Items)
}
// PlayerActions définit les actions nécessaires pour utiliser une potion.
type PlayerActions interface {
	Heal(amount int)
	TakeDamage(amount int)
}

// UseItem utilise une potion de l'inventaire.
func (i *Inventory) UseItem(index int, player PlayerActions) bool {
	if index < 0 || index >= len(i.Items) {
		return false
	}

	item := i.Items[index]

	switch item.Name {
	case HealingPotion.Name:
		player.Heal(item.Value)

	case PoisonPotion.Name:
		player.TakeDamage(item.Value)

	default:
		return false
	}

	// Retire la potion après utilisation.
	i.RemoveItem(index)

	return true
}

