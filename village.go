package main

import "fmt"

func Village() {
	fmt.Println("VILLAGE DE MARCHANG")
	fmt.Println("1. Aller aux prairies")
	fmt.Println("2. Aller dans la forêt")
	fmt.Println("3. Aller chez le forgeron")

	var choix int
	fmt.Scan(&choix)

	if choix == 1 {
	fmt.Println("Vous allez dans les prairies.")
}

    if choix == 2 {
	fmt.Println("Vous allez dans la forêt.")
}

    if choix == 3 {
	fmt.Println("Vous allez chez le forgeron.")
}
}
