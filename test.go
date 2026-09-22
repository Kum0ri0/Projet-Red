package main

import (
	"fmt"
)

func main() {

	// Création du personnage
	rudeus := CreatePlayer("Rudeus", "Humain")

	fmt.Println("=================================")
	fmt.Println("      CRÉATION DU PERSONNAGE")
	fmt.Println("=================================")

	fmt.Println("Nom :", rudeus.Name)
	fmt.Println("Classe :", rudeus.Class)
	fmt.Printf("PV : %d/%d\n", rudeus.HP, rudeus.MaxHP)
	fmt.Println("Niveau :", rudeus.Level)
	fmt.Printf("XP : %d/%d\n", rudeus.XP, rudeus.XPToNextLevel())
	fmt.Println("Argent :", rudeus.Gold)

	// Test des dégâts
	fmt.Println()
	fmt.Println("Rudeus reçoit 30 dégâts...")
	rudeus.TakeDamage(30)

	fmt.Printf("PV : %d/%d\n", rudeus.HP, rudeus.MaxHP)

	// Test du soin
	fmt.Println()
	fmt.Println("Rudeus récupère 10 PV...")
	rudeus.Heal(10)

	fmt.Printf("PV : %d/%d\n", rudeus.HP, rudeus.MaxHP)

	// Test de l'XP
	fmt.Println()
	fmt.Println("Rudeus gagne 100 XP...")
	rudeus.GainXP(100)

	fmt.Printf("Niveau : %d\n", rudeus.Level)
	fmt.Printf("PV : %d/%d\n", rudeus.HP, rudeus.MaxHP)
	fmt.Printf("XP : %d/%d\n", rudeus.XP, rudeus.XPToNextLevel())

	// Test de l'argent
	fmt.Println()
	fmt.Println("Rudeus gagne 50 pièces d'or...")
	rudeus.AddGold(50)

	fmt.Println("Argent :", rudeus.Gold)

	// ==============================
	// TEST DES ITEMS
	// ==============================

	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("          TEST DES ITEMS")
	fmt.Println("=================================")

	// Potion de soin
	fmt.Println()
	fmt.Println("Nom :", HealingPotion.Name)
	fmt.Println("Type :", HealingPotion.Type)
	fmt.Println("Description :", HealingPotion.Description)

	// Potion de poison
	fmt.Println()
	fmt.Println("Nom :", PoisonPotion.Name)
	fmt.Println("Type :", PoisonPotion.Type)
	fmt.Println("Description :", PoisonPotion.Description)

	// Hache
	fmt.Println()
	fmt.Println("Nom :", Axe.Name)
	fmt.Println("Type :", Axe.Type)
	fmt.Println("Description :", Axe.Description)

	// Épée
	fmt.Println()
	fmt.Println("Nom :", Sword.Name)
	fmt.Println("Type :", Sword.Type)
	fmt.Println("Description :", Sword.Description)

	// Bâton de mage
	fmt.Println()
	fmt.Println("Nom :", MageStaff.Name)
	fmt.Println("Type :", MageStaff.Type)
	fmt.Println("Description :", MageStaff.Description)

	// Livre de sorts
	fmt.Println()
	fmt.Println("Nom :", Spellbook.Name)
	fmt.Println("Type :", Spellbook.Type)
	fmt.Println("Description :", Spellbook.Description)

	// Fourrure
	fmt.Println()
	fmt.Println("Nom :", Fur.Name)
	fmt.Println("Type :", Fur.Type)
	fmt.Println("Description :", Fur.Description)

	// Cuir
	fmt.Println()
	fmt.Println("Nom :", Leather.Name)
	fmt.Println("Type :", Leather.Type)
	fmt.Println("Description :", Leather.Description)

	// Peau
	fmt.Println()
	fmt.Println("Nom :", Hide.Name)
	fmt.Println("Type :", Hide.Type)
	fmt.Println("Description :", Hide.Description)

	// Bois
	fmt.Println()
	fmt.Println("Nom :", Wood.Name)
	fmt.Println("Type :", Wood.Type)
	fmt.Println("Description :", Wood.Description)

	// Métal
	fmt.Println()
	fmt.Println("Nom :", Metal.Name)
	fmt.Println("Type :", Metal.Type)
	fmt.Println("Description :", Metal.Description)

	// ==============================
	// TEST DE L'INVENTAIRE
	// ==============================

	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("        TEST DE L'INVENTAIRE")
	fmt.Println("=================================")

	// Ajouter des objets
	rudeus.Inventory.AddItem(HealingPotion)
	rudeus.Inventory.AddItem(PoisonPotion)
	rudeus.Inventory.AddItem(Sword)

	fmt.Println("Nombre d'objets :", rudeus.Inventory.Size())

	// Afficher les objets
	for index, item := range rudeus.Inventory.Items {
		fmt.Printf("[%d] %s\n", index, item.Name)
	}

	// Utiliser la potion de soin
	fmt.Println()
	fmt.Println("Rudeus utilise une potion de soin...")

	rudeus.TakeDamage(60)

	fmt.Printf("PV avant la potion : %d/%d\n", rudeus.HP, rudeus.MaxHP)

	rudeus.Inventory.UseItem(0, &rudeus)

	fmt.Printf("PV après la potion : %d/%d\n", rudeus.HP, rudeus.MaxHP)

	// Utiliser la potion de poison
	fmt.Println()
	fmt.Println("Rudeus utilise une potion de poison...")

	rudeus.Inventory.UseItem(0, &rudeus)

	fmt.Printf("PV après le poison : %d/%d\n", rudeus.HP, rudeus.MaxHP)

	// Afficher l'inventaire restant
	fmt.Println()
	fmt.Println("Inventaire restant :")

	for index, item := range rudeus.Inventory.Items {
		fmt.Printf("[%d] %s\n", index, item.Name)
	}
}
