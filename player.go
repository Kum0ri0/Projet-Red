package main

type Player struct {
	Name      string
	Class     string
	HP        int
	MaxHP     int
	Level     int
	XP        int
	Gold      int
	Inventory Inventory
}

// CreatePlayer crée un nouveau personnage avec les valeurs de départ.
func CreatePlayer(name string, class string) Player {
	return Player{
		Name:  name,
		Class: class,
		HP:    100,
		MaxHP: 100,
		Level: 1,
		XP:    0,
		Gold:  0,
	}
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
func (p *Player) XPToNextLevel() int {
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