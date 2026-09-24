package main

import (
	"fmt"
)	

func Village(joueur *Player, inv *Inventory) {
	for {
		fmt.Println("====== VILLAGE DE MARCHANG ======")
		fmt.Println("")
		fmt.Println("[1] Aller aux", "prairies", "🌿")
		fmt.Println("[2] Aller dans la","forêt", "🌲")
		fmt.Println("[3] Aller chez le ", "forgeron", "⚒️")
		fmt.Println("[4] Voir l'inventaire", "🎒")
		fmt.Println("[5] Aller à la mine", "⛏️")
		fmt.Println("[6] Aller chez le marchand", "🏪")
		fmt.Println("[7] Aller chez le médecin", "🏥")
		fmt.Println("[8] Voir les quêtes", "📜")
		fmt.Println("[9] Quitter le Village")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			Prairies(inv)
		}

		if choix == 2 {
			Foret(inv)
		}

		if choix == 3 {
			Forgeron(inv)
		}

		if choix == 4 {
			inv.Display()
		}

		if choix == 5 {
			Mine(inv)
		}

		if choix == 6 {
			Marchand(joueur, inv)
		}

		if choix == 7 {
			Medecin(joueur)
		}

		if choix == 8 {
			Quetes(joueur)
		}

		if choix == 9 {
			return
		}
	}
}

