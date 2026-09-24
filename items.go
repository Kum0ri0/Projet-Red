package main

import "fmt"

type Item struct {
	Name        string
	Type        string
	Description string
	Value       int
	Spell		Spell
}

const (
	TypeWeapon   = "Arme"
	TypePotion   = "Potion"
	TypeResource = "Ressource"
	TypeBook = "Livre"
)

// =========================
// ARMES
// =========================

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
var DragonSword = Item{
	Name:        "épée du Dragon",
	Type:        TypeWeapon,
	Description: "L'épée obtenue aprés avoir vaincu le Dragon",
}

var Dagger = Item{
	Name:        "Dague",
	Type:        TypeWeapon,
	Description: "Une dague",
}

var MageStaff = Item{
	Name:        "Bâton de mage",
	Type:        TypeWeapon,
	Description: "Un bâton utilisé par les mages.",
	Value:       30,
}

var WoodenSword = Item{
	Name:        "Épée en bois",
	Type:        TypeWeapon,
	Description: "Une épée fabriquée en bois.",
	Value:       20,
}

var MetalSword = Item{
	Name:        "Épée en métal",
	Type:        TypeWeapon,
	Description: "Une épée fabriquée en métal.",
	Value:       40,
}

// =========================
// POTIONS
// =========================

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

// Ajout de livre pour les sorts 
// =========================
// books
// =========================
var LivreFireball = Item{
    Name: "Grimoire : Fireball",
    Type: TypeBook,
    Spell: Fireball,
	Value: 30,
}

var LivreStorm = Item{
    Name: "Grimoire : Tempète",
    Type: TypeBook,
    Spell: Storm,
}

var LivreDivinelight = Item{
    Name: "Grimoire : Lumière Divine",
    Type: TypeBook,
    Spell: Divinelight,
}

var LivrewaterBlade = Item{
    Name: "Grimoire : lame d'eau",
    Type: TypeBook,
    Spell: waterBlade,
	Value: 40,
}

var LivreDragonBreath = Item{
    Name: "Grimoire : Souffle du Dragon",
    Type: TypeBook,
    Spell: DragonBreath,
}



var LivreBlizzard = Item{
    Name: "Grimoire : blizarre",
    Type: TypeBook,
    Spell: Blizzard,
}

// =========================
// RESSOURCES
// =========================

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

// =========================
// AFFICHAGE
// =========================

func (i Item) Display() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║            OBJET             ║")
	fmt.Println("╠══════════════════════════════╣")
	fmt.Printf("║ Nom         : %-14s ║\n", i.Name)
	fmt.Printf("║ Type        : %-14s ║\n", i.Type)
	fmt.Printf("║ Description : %-14s ║\n", i.Description)
	fmt.Printf("║ Valeur      : %-14d ║\n", i.Value)
	fmt.Println("╚══════════════════════════════╝")
}

func DisplayAll() {
	items := []Item{
		Axe,
		Sword,
		MageStaff,
		WoodenSword,
		MetalSword,
		HealingPotion,
		PoisonPotion,
		Fur,
		Leather,
		Hide,
		Wood,
		Metal,
	}

	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║          OBJETS              ║")
	fmt.Println("╚══════════════════════════════╝")

	for _, item := range items {
		fmt.Printf("- %-20s [%s]\n", item.Name, item.Type)
	}
}
