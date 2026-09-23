package main

import (
	"fmt"
	"math/rand"
)



type Monster struct {
	emoji  string
	nom    string
	pvMax  int
	pv     int
	degats int
	xp     int
	orMin  int
	orMax  int
	drops  []Drop
	attaques []Attaque
}

type Drop struct {
	emoji  string
	item    Item
	taux   int
	nbrMin int
	nbrMax int
}

type Attaque struct {
	nom string 
	degats int
	precision int 
}



func initSlime() Monster {
	return Monster{
		emoji:  "🧊",
		nom:    "Slime",
		pvMax:  20,
		pv:     20,
		degats: 2,
		xp:     10,
		orMin:  1,
		orMax:  3,
		attaques: []Attaque{
			{nom: "Bloup", degats: 5, precision: 95},  
		},
	}
}

func initGoblin() Monster {
	return Monster{
		emoji:  "🧌",
		nom:    "Gobelin",
		pvMax:  50,
		pv:     50,
		degats: 5,
		xp:     20,
		orMin:  4,
		orMax:  7,
		drops: []Drop{
			{emoji: "🗡️", item: Dagger, taux: 2, nbrMin: 1, nbrMax: 1},
		},
		attaques: []Attaque{
			{nom: "Coups de poing", degats: 10, precision: 70}, 
			{nom: "Entaille", degats: 5, precision: 95},  
		},
		
	}
}

func initLoup() Monster {
	return Monster{
		emoji:  "🐺",
		nom:    "Loup",
		pvMax:  70,
		pv:     70,
		degats: 10,
		xp:     30,
		orMin:  6,
		orMax:  8,
		drops: []Drop{
			{emoji: "🧶", item : Fur , taux: 70, nbrMin: 1, nbrMax: 2},
		},
		attaques: []Attaque{
			{nom: "Morsure", degats: 10, precision: 70}, 
			{nom: "Griffure", degats: 5, precision: 95}, 
			{nom: "Assaut", degats: 20, precision: 40}, 
		},
	}
}

func initSanglier() Monster {
	return Monster{
		emoji:  "🐗",
		nom:    "Sanglier",
		pvMax:  90,
		pv:     90,
		degats: 7,
		xp:     20,
		orMin:  10,
		orMax:  13,
		drops: []Drop{
			{emoji: "🧥", item: Leather , taux: 40, nbrMin: 1, nbrMax: 3},
		},
		attaques: []Attaque{
			{nom: "Morsure", degats: 12, precision: 70}, 
			{nom: "Charge", degats: 7, precision: 95}, 
			{nom: "Assaut", degats: 21, precision: 40}, 
		},
	}
}

func initMage() Monster {
	return Monster{
		emoji:  "🧙",
		nom:    "Mage",
		pvMax:  100,
		pv:     100,
		degats: 15,
		xp:     40,
		orMin:  14,
		orMax:  17,
		drops: []Drop{
			{emoji: "🪄", item: MageStaff, taux: 40, nbrMin: 1, nbrMax: 1},
		},
		attaques: []Attaque{
			{nom: "Boule de feu", degats: 13, precision: 70}, 
			{nom: "Éclair", degats: 9, precision: 95}, 
			{nom: "Météore", degats: 26, precision: 40}, 
		},
	}
}

func initOgre() Monster {
	return Monster{
		emoji:  "👹",
		nom:    "Ogre",
		pvMax:  130,
		pv:     130,
		degats: 20,
		xp:     60,
		orMin:  18,
		orMax:  21,
		drops: []Drop{
			{emoji: "🪨", item: Hide, taux: 30, nbrMin: 1, nbrMax: 1},
		},
		attaques: []Attaque{
			{nom: "Coups de massue", degats: 14, precision: 70}, 
			{nom: "Frappe foudroyante", degats: 10, precision: 95}, 
			{nom: "Assaut", degats: 30, precision: 40}, 
		},
	}
}

func initDragon() Monster {
	return Monster{
		emoji:  "🐉",
		nom:    "Dragon",
		pvMax:  200,
		pv:     200,
		degats: 30,
		xp:     100,
		orMin:  30,
		orMax:  35,
		drops: []Drop{
			{emoji: "⚔️", item: DragonSword, taux: 5, nbrMin: 1, nbrMax: 1},
		},
		attaques: []Attaque{
			{nom: "Lance flamme", degats: 50, precision: 70}, 
			{nom: "Coups d'aile", degats: 30, precision: 95}, 
			{nom: "Ruée", degats: 70, precision: 40}, 
		},
	}
}

func displayMonster(m Monster) {
	fmt.Println(m.emoji, "|", m.nom, ":")
	fmt.Println()
	fmt.Println(pvBar(m.pv, m.pvMax), m.pv, "/", m.pvMax, "pv")
}

func damageMonster(m *Monster, dmg int) {
	m.pv -= dmg
	if m.pv < 0 {
		m.pv = 0
	}
}

func isMonsterDead(m Monster) bool {
	if m.pv <= 0 {
		return true
	}
	return false
}

func rollDrops(m Monster, p *Player) {
	for _, r := range m.drops {
		if r.taux > rand.Intn(100) {
			resultatDrop := rand.Intn(r.nbrMax-r.nbrMin+1) + r.nbrMin
 
			for i := 0; i < resultatDrop; i++ {
				if !p.Inventory.AddItem(r.item) {
					fmt.Println("Inventaire plein, tu perds", r.item.Name)
					break
				}
			}
 
			fmt.Println("Bien joué, le", m.nom, "a laissé", resultatDrop, r.emoji, r.item.Name)
		}
	}
}

func rollGold(m Monster, p *Player) {
	resultatOr := rand.Intn(m.orMax-m.orMin+1) + m.orMin
	p.AddGold(resultatOr)
	fmt.Println("Le", m.nom, "a laissé", resultatOr, "pièces d'or")
}