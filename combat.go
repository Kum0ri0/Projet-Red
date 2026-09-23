package main

import (
	"fmt"
	"strings"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func displayPlayer(p Player) {
	fmt.Println("|", p.Name, ":")
	fmt.Println()
	fmt.Println(pvBar(p.HP, p.MaxHP), p.HP, "/", p.MaxHP, "pv")
}

func monsterAttack(m Monster, p *Player, tour int) {
	numero := rand.Intn(len(m.attaques))
	attaque := m.attaques[numero]

	if  attaque.precision > rand.Intn(100) {
		degats := attaque.degats

		if tour % 3 == 0  {
			degats = degats*2
			fmt.Println("🔥", m.nom, "est enragé ! Dégâts doublés !")
		}

		p.TakeDamage(degats)
		fmt.Println("Le", m.nom, "utilise", attaque.nom, "et inflige", degats, "dégâts à", p.Name)
	} else {
		fmt.Println("Le", m.nom, "utilise", attaque.nom, "... mais rate son attaque !")
	}
}

	
func ChooseSpell(m *Monster, p *Player) bool {
	spells := sortsDuJoueur()

	for i, s := range spells {
		fmt.Printf("[%d] %s (%d dégâts)\n", i+1, s.nom, s.degats)
	}
	fmt.Println("[0] Retour")
 
	var choix int
	fmt.Scanln(&choix)
 
	if choix == 0 {
		return false
	}
	if choix < 1 || choix > len(spells) {
		fmt.Println("Ce choix n'existe pas")
		return false
	}
 
	spell := spells[choix-1]
 
	damageMonster(m, spell.degats)
 
	fmt.Println()
	fmt.Println(p.Name, "lance", spell.nom, "et inflige", spell.degats, "dégâts au", m.nom)
	fmt.Println()
 
	return true
}

func playerAttack(m *Monster, p Player) {
	attaqueBasique := 5
	damageMonster(m, attaqueBasique)

	fmt.Println()
	fmt.Println(p.Name, "a infligé", attaqueBasique, "dégâts au", m.nom, "avec l'attaque basique")
	fmt.Println()
}

func pvBar(pv int, pvMax int) string {
	pleine := pv * 20 / pvMax
	vide := 20 - pleine

	return strings.Repeat("█", pleine) + strings.Repeat("░", vide)
}

func characterTurn(m *Monster, p *Player) (int, bool) {
	var choix int
	fmt.Scanln(&choix)
 
	switch choix {
	case 1:
		return choix, ChooseSpell(m, p)
	case 2:
		return choix, useInventory(p)
	case 3:
		return choix, true
	default:
		fmt.Println("Choix invalide")
		return choix, false
	}
}

func displayCombat(m Monster, p Player, tour int) {
	fmt.Println("========== TOUR", tour, "==========")
	fmt.Println()
	fmt.Println(strings.Repeat(" ", 20), m.emoji, m.nom, ":")
	fmt.Println()
	fmt.Println(strings.Repeat(" ", 20), pvBar(m.pv, m.pvMax), m.pv, "/", m.pvMax, "pv")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(strings.Repeat(" ", 5), p.Name, ":")
	fmt.Println()
	fmt.Println(strings.Repeat(" ", 5), pvBar(p.HP, p.MaxHP), p.HP, "/", p.MaxHP, "pv")
	fmt.Println()
	fmt.Println()
	fmt.Println("[1] Attaque ", strings.Repeat(" ", 10), "[2] Objets ")
	fmt.Println()
	fmt.Println("[3] Fuir ")
}

func trainingFight(m *Monster, p *Player) {
	tour := 1

	for {
		displayCombat(*m, *p, tour)

		choix, aJoue := characterTurn(m, p)

		if choix == 3 {
			break
		}
		if !aJoue {
			time.Sleep(1 * time.Second)
			clearScreen()
			continue
		}

		
		if isMonsterDead(*m) {
			fmt.Println("====== Victoire ! 🎉 ======")

			rollGold(*m, p)

			p.GainXP(m.xp)
			fmt.Println("✨", p.Name, "gagne", m.xp, "XP")

			rollDrops(*m, p)
			break
		}

		monsterAttack(*m, p, tour)

		if p.HP <= 0 {
			fmt.Println("===== Perdu =====")
			fmt.Println("Vous êtes mort")
			break
		}
		time.Sleep(2 * time.Second)
		clearScreen()
		tour++
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}


func useInventory(p *Player) bool {
	if p.Inventory.Size() == 0 {
		fmt.Println("Ton inventaire est vide ..")
		return false
	}
 
	for i, r := range p.Inventory.Items {
		fmt.Printf("[%d] %s \n", i+1, r.Name)
	}
	fmt.Println("[0] Retour")
 
	var choix int
	fmt.Scanln(&choix)
 
	if choix == 0 {
		return false
	}
	if choix < 1 || choix > len(p.Inventory.Items) {
		fmt.Println("Ce choix n'existe pas")
		return false
	}
 
	item := p.Inventory.Items[choix-1]
	fmt.Println("Tu utilises", item.Name)
	p.Inventory.UseItem(choix-1, p)
 
	return true
}

