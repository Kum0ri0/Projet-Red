package main

func main() {
	perso := CreatePlayer("Subaru", "Humain")

	perso.Inventory.AddItem(HealingPotion)
	perso.Inventory.AddItem(HealingPotion)
	perso.Inventory.AddItem(PoisonPotion)

	useInventory(&perso)
}