package main

import "fmt"

type Item struct {
	Name        string
	Type        string
	Description string
	Value       int
}

// Types d'objets.
const (
	TypeWeapon   = "Arme"
	TypePotion   = "Potion"
	TypeResource = "Ressource"
)

// Armes
var Axe = Item{
	Name:        "Hache",
	Type:        TypeWeapon,
	Description: "Une hache solide.",
	Value:       20,
}

var Sword = Item{
	Name:        "Épée",
	Type:        TypeWeapon,
	Description: "Une épée tranchante.",
	Value:       25,
}

var MageStaff = Item{
	Name:        "Bâton de mage",
	Type:        TypeWeapon,
	Description: "Un bâton utilisé par les mages.",
	Value:       30,
}

// Potions
var HealingPotion = Item{
	Name:        "Potion de soin",
	Type:        TypePotion,
	Description: "Restaure 50 PV.",
	Value:       50,
}

var PoisonPotion = Item{
	Name:        "Potion de poison",
	Type:        TypePotion,
	Description: "Fait perdre 30 PV.",
	Value:       30,
}

// Ressources
var Fur = Item{
	Name:        "Fourrure",
	Type:        TypeResource,
	Description: "Une fourrure récupérée sur un animal.",
	Value:       5,
}

var Leather = Item{
	Name:        "Cuir",
	Type:        TypeResource,
	Description: "Du cuir pouvant servir à fabriquer des objets.",
	Value:       5,
}

var Hide = Item{
	Name:        "Peau",
	Type:        TypeResource,
	Description: "Une peau récupérée sur un animal.",
	Value:       4,
}

var Wood = Item{
	Name:        "Bois",
	Type:        TypeResource,
	Description: "Un morceau de bois.",
	Value:       2,
}

var Metal = Item{
	Name:        "Métal",
	Type:        TypeResource,
	Description: "Un morceau de métal.",
	Value:       8,
}
var Dagger = Item{
	Name:        "Dague",
	Type:        TypeWeapon,
	Description: "Une Dague",
}
var DragonSword = Item{
	Name:        "épée du dragon",
	Type:        TypeWeapon,
	Description: "Une épée utilisant la puissance des dragons",
}

// Display affiche les informations d'un objet.
func (i Item) Display() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║             OBJET            ║")
	fmt.Println("╠══════════════════════════════╣")
	fmt.Printf("║ Nom         : %-15s ║\n", i.Name)
	fmt.Printf("║ Type        : %-15s ║\n", i.Type)
	fmt.Printf("║ Description : %-15s ║\n", i.Description)
	fmt.Printf("║ Valeur      : %-15d ║\n", i.Value)
	fmt.Println("╚══════════════════════════════╝")
}

// DisplayAll affiche tous les objets disponibles.
func DisplayAll() {
	items := []Item{
		Axe,
		Sword,
		MageStaff,
		HealingPotion,
		PoisonPotion,
		Fur,
		Leather,
		Hide,
		Wood,
		Metal,
	}

	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║       OBJETS DISPONIBLES     ║")
	fmt.Println("╠══════════════════════════════╣")

	for _, item := range items {
		fmt.Printf("║ %-20s ║\n", item.Name)
	}

	fmt.Println("╚══════════════════════════════╝")
}
