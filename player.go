package main

import "fmt"

type Spell struct {
	Name   string
	Damage int
}

var Punch = Spell{
	Name:   "Coup de poing",
	Damage: 8,
}

var Fireball = Spell{
	Name:   "Boule de Feu",
	Damage: 18,
}
// ajout ========================================
var Storm = Spell{
	Name:   "tempête",
	Damage: 20,
}

var Divinelight = Spell{
	Name:   "Lumière Divine",
	Damage: 25,
}

var waterBlade = Spell{
	Name:   "Lame d'eau",
	Damage: 16,
}

var DragonBreath = Spell{
	Name:   "Souffle du Dragon",
	Damage: 80,
}

var Blizzard = Spell{
	Name: "Blizzard",
	Damage: 30,
}
// ajout ========================================
type Player struct {
	Name      string
	Class     string
	HP        int
	MaxHP     int
	Level     int
	XP        int
	Gold      int
	Spells    []Spell
	Inventory Inventory
}

// CreatePlayer crée un nouveau personnage avec Coup de poing.
func CreatePlayer(name string, class string) Player {
	return Player{
		Name:   name,
		Class:  class,
		HP:     100,
		MaxHP:  100,
		Level:  1,
		XP:     0,
		Gold:   0,
		Spells: []Spell{Punch},
	}
}
// LearnSpell ajoute un sort au joueur s'il ne le connaît pas déjà.
// Retourne true si le sort a été appris, false s'il était déjà connu.
func (p *Player) LearnSpell(s Spell) bool {
	for _, spell := range p.Spells {
		if spell.Name == s.Name {
			return false
		}
	}

	p.Spells = append(p.Spells, s)
	return true
}


// Display affiche les informations du joueur.
func (p Player) Display() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║          PERSONNAGE          ║")
	fmt.Println("╠══════════════════════════════╣")
	fmt.Printf("║ Nom      : %-18s ║\n", p.Name)
	fmt.Printf("║ Classe   : %-18s ║\n", p.Class)
	fmt.Printf("║ PV       : %-18s ║\n", fmt.Sprintf("%d / %d", p.HP, p.MaxHP))
	fmt.Printf("║ Niveau   : %-18d ║\n", p.Level)
	fmt.Printf("║ XP       : %-18d ║\n", p.XP)
	fmt.Printf("║ Or       : %-18d ║\n", p.Gold)
	fmt.Println("╚══════════════════════════════╝")
}

// TakeDamage retire des PV au joueur.
func (p *Player) TakeDamage(damage int) {
	p.HP -= damage

	if p.HP < 0 {
		p.HP = 0
	}
}

// Heal soigne le joueur sans dépasser ses PV maximum.
func (p *Player) Heal(amount int) {
	p.HP += amount

	if p.HP > p.MaxHP {
		p.HP = p.MaxHP
	}
}

// GainXP ajoute de l'expérience au joueur.
func (p *Player) GainXP(amount int) {
	p.XP += amount

	for p.XP >= p.XPToNextLevel() {
		p.XP -= p.XPToNextLevel()
		p.LevelUp()
	}
}

// XPToNextLevel retourne l'XP nécessaire pour passer au niveau suivant.
func (p Player) XPToNextLevel() int {
	return p.Level * 100
}

// LevelUp fait monter le joueur d'un niveau.
func (p *Player) LevelUp() {
	p.Level++
	p.MaxHP += 20
	p.HP = p.MaxHP
}

// AddGold ajoute de l'argent au joueur.
func (p *Player) AddGold(amount int) {
	p.Gold += amount
}

// RemoveGold retire de l'argent au joueur.
func (p *Player) RemoveGold(amount int) bool {
	if amount > p.Gold {
		return false
	}

	p.Gold -= amount
	return true
}
