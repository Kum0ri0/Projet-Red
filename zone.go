package main 

import ("fmt"
		"math/rand"
		"time"

)


type Zone struct {
	emoji string 
	nom string 
	niveauMin int
	mobs []Monster 
}

func initPrairie() Zone {
	return Zone {
		emoji: "🌿",
		nom: "Prairie calme",
		niveauMin: 1,
		mobs: []Monster{initSlime(), initGoblin()},
}
}

func initForet() Zone {
	return Zone {
		emoji: "🌲 ",
		nom: "Forêt sombre",
		niveauMin: 3,
		mobs: []Monster{initSanglier(), initLoup()}, 
}}

func initMontagne() Zone{
	return Zone {
		emoji: "⛰️ ",
		nom: "Montagne maudite",
		niveauMin: 5,
		mobs: []Monster{initMage(), initOgre()}, 
}
}
func initAntre() Zone {
	return Zone {
		emoji: "🌋 ",
		nom: "Antre du Dragon",
		niveauMin: 10,
		mobs: []Monster{initDragon()}, 
}
}

func allZones() []Zone  {
	return []Zone{initPrairie(), initForet(), initMontagne(), initAntre()}
}

func EntreZone(z Zone, m Monster) {
	fmt.Printf("Vous entrez dans %s %s.... \n",  z.emoji, z.nom)
	fmt.Println()
	fmt.Printf("Un %s %s Sauvage vous attaque ! \n",  m.emoji, m.nom)
	time.Sleep(2 * time.Second)
	clearScreen()

}
func AfficheZone(p Player) {
fmt.Println("=== Où veux-tu aller ? ===")
fmt.Println("")
for i, z := range allZones() {
	if p.Level >= z.niveauMin {
		fmt.Printf("[%d] %s %s (niv. %d)\n", i+1, z.emoji, z.nom, z.niveauMin)
	}else {
		fmt.Printf("[%d] %s %s (niv. %d) 🔒\n", i+1, z.emoji, z.nom, z.niveauMin)
		
	}
}
}


func chooseZone(p *Player) {
	AfficheZone(*p)
	var choix int
	fmt.Scanln(&choix)

	zones := allZones()
	if choix < 1 || choix > len(zones) {
		fmt.Println("Ce choix n'existe pas")
		return 
	}
	zone := zones[choix-1]
	if p.Level < zone.niveauMin {
	   fmt.Println("Cette zone est trop dangereuse pour toi !")
	   return
	}
	numero := rand.Intn(len(zone.mobs))
	mob := zone.mobs[numero]
	EntreZone(zone, mob)

	trainingFight(&mob, p)

}

