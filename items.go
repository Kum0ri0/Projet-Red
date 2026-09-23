package main

type Item struct {
	Name        string
	Type        string
	Description string
	Value 	int 
}
const (
	TypePotion    = "Potion"
	TypeWeapon    = "Arme"
	TypeSpellbook = "Livre de sorts"
	TypeMaterial  = "Matériau"
)
var HealingPotion = Item{
	Name:        "Potion de soin",
	Type:        TypePotion,
	Description: "Restaure des PV",
	Value: 50,
}

var PoisonPotion = Item{
	Name:        "Potion de poison",
	Type:        TypePotion,
	Description: "Une potion empoisonnée",
	Value : 30, 
}

var Axe = Item{
	Name:        "Hache",
	Type:        TypeWeapon,
	Description: "Une hache",
}

var Dagger = Item{
	Name:        "Dague",
	Type:        TypeWeapon,
	Description: "Une Dague",
}

var Sword = Item{
	Name:        "Épée",
	Type:        TypeWeapon,
	Description: "Une épée",
}

var MageStaff = Item{
	Name:        "Bâton de mage",
	Type:        TypeWeapon,
	Description: "Un bâton utilisé par les mages",
}

var DragonSword = Item{
	Name:        "épée du dragon",
	Type:        TypeWeapon,
	Description: "Une épée utilisant la puissance des dragons",
}

var Spellbook = Item{
	Name:        "Livre de sorts",
	Type:        TypeSpellbook,
	Description: "Un livre contenant des sorts",
}
var Fur = Item{
	Name:        "Fourrure",
	Type:        TypeMaterial,
	Description: "Une fourrure utilisée comme matériau",
}

var Leather = Item{
	Name:        "Cuir",
	Type:        TypeMaterial,
	Description: "Du cuir utilisé comme matériau",
}

var Hide = Item{
	Name:        "Peau",
	Type:        TypeMaterial,
	Description: "Une peau utilisée comme matériau",
}

var Wood = Item{
	Name:        "Bois",
	Type:        TypeMaterial,
	Description: "Du bois utilisé comme matériau",
}

var Metal = Item{
	Name:        "Métal",
	Type:        TypeMaterial,
	Description: "Du métal utilisé comme matériau",
}