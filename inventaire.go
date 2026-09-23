package main

import "fmt"

var bois int
var metal int
var epee int
var epeeBois int
var epeeMetal int
var argent int = 100
var vie int = 100
var vieMax int = 100

func Inventaire() {
	fmt.Println("INVENTAIRE")
	fmt.Println("Bois :", bois)
	fmt.Println("Métal :", metal)
	fmt.Println("Épées :", epee)
	fmt.Println("Épées en bois :", epeeBois)
    fmt.Println("Épées en métal :", epeeMetal)
	fmt.Println("Argent :", argent)
	fmt.Println("Vie :", vie, "/", vieMax)
}
