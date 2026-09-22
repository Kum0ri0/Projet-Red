package main

import "fmt"

func Village() {
	for {
		fmt.Println("VILLAGE DE MARCHANG")
		fmt.Println("1. Aller aux prairies")
		fmt.Println("2. Aller dans la forêt")
		fmt.Println("3. Aller chez le forgeron")
		fmt.Println("4. Voir l'inventaire")
		fmt.Println("5. Quitter le jeu")

		var choix int
		fmt.Scan(&choix)

		if choix == 1 {
			Prairies()
		}

		if choix == 2 {
			fmt.Println("Vous allez dans la forêt.")
		}
        
		if choix == 3 {
	         fmt.Println("5. Aller à la mine")
		}
			
		if choix == 4 {
			Forgeron()
		}

		if choix == 5 {
			Inventaire()
		}

		if choix == 6 {
			return
		}
	}
}
