package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	clearScreen()
	afficherTitre()

	joueur := creerPersonnage()

	clearScreen()
	fmt.Println("Bienvenue,", joueur.Name, "le", joueur.Class, "!")
	fmt.Println("Ton aventure commence au village de Marchang...")
	fmt.Println()
	joueur.Display()
	attendreEntree()

	menuPrincipal(&joueur)

	clearScreen()
	fmt.Println("Merci d'avoir joué,", joueur.Name, "! À bientôt 👋")
}

// =========================
// ÉCRAN TITRE
// =========================

func afficherTitre() {
	fmt.Println("╔══════════════════════════════╗")
	fmt.Println("║          PROJET RED          ║")
	fmt.Println("║   Les aventures de Marchang  ║")
	fmt.Println("╚══════════════════════════════╝")
	fmt.Println()
}

// =========================
// CRÉATION DU PERSONNAGE
// =========================

func creerPersonnage() Player {
	var nom string
	for nom == "" {
		fmt.Print("Quel est ton nom, aventurier ? ")
		nom = lireLigne()
	}

	fmt.Println()
	fmt.Println("Choisis ta classe :")
	fmt.Println("[1] Guerrier  (+20 PV max)")
	fmt.Println("[2] Mage      (connaît Boule de Feu)")
	fmt.Println("[3] Voleur    (commence avec 30 pièces d'or)")

	classe := lireChoix(1, 3)

	return nouveauPersonnage(nom, classe)
}

// nouveauPersonnage crée le joueur selon la classe choisie
// (1 = Guerrier, 2 = Mage, 3 = Voleur). Utilisé par les deux versions du jeu.
func nouveauPersonnage(nom string, classe int) Player {
	var joueur Player
	switch classe {
	case 1:
		joueur = CreatePlayer(nom, "Guerrier")
		joueur.MaxHP += 20
		joueur.HP = joueur.MaxHP
	case 2:
		joueur = CreatePlayer(nom, "Mage")
		joueur.LearnSpell(Fireball)
	case 3:
		joueur = CreatePlayer(nom, "Voleur")
		joueur.AddGold(30)
	}

	// Kit de départ commun à toutes les classes.
	joueur.Inventory.AddItem(HealingPotion)
	joueur.Inventory.AddItem(HealingPotion)
	joueur.AddGold(50) // tempo
	return joueur
}

// =========================
// MENU PRINCIPAL
// =========================

func menuPrincipal(joueur *Player) {
	for {
		clearScreen()
		afficherFicheJoueur(joueur)
		fmt.Println()
		fmt.Printf("  [1] %-13s %s\n", "Village", "Récolter, forger, acheter")
		fmt.Printf("  [2] %-13s %s\n", "Aventure", "Partir combattre des monstres")
		fmt.Printf("  [3] %-13s %s\n", "Personnage", "Voir mes stats et mon sac")
		fmt.Printf("  [4] %-13s %s\n", "Objets", "Boire une potion")
		fmt.Printf("  [5] %s\n", "Quitter")
		fmt.Println()

		switch lireChoix(1, 5) {
		case 1:
			clearScreen()
			// Le village utilise l'inventaire du joueur : ce qui est récolté
			// ou acheté est aussi disponible en combat, et inversement.
			Village(joueur, &joueur.Inventory)
		case 2:
			clearScreen()
			chooseZone(joueur)
			if joueur.HP <= 0 {
				reanimer(joueur)
			}
			attendreEntree()
		case 3:
			clearScreen()
			joueur.Display()
			fmt.Printf("XP : %d / %d\n", joueur.XP, joueur.XPToNextLevel())
			fmt.Println()
			joueur.Inventory.Display()
			attendreEntree()
		case 4:
			menuObjets(joueur)
		case 5:
			return
		}
	}
}

// largeurFiche est la largeur intérieure du cadre du menu principal.
const largeurFiche = 46

// afficherFicheJoueur affiche le cadre du menu principal : titre, stats du joueur
// et prochaine zone à débloquer.
func afficherFicheJoueur(joueur *Player) {
	bord := strings.Repeat("═", largeurFiche)

	fmt.Println("╔" + bord + "╗")
	ligneFiche(centrer("PROJET  RED"))
	ligneFiche(centrer("Le village de Marchang"))
	fmt.Println("╠" + bord + "╣")
	ligneFiche(deuxCotes("  "+joueur.Name+" le "+joueur.Class, fmt.Sprintf("Niveau %d  ", joueur.Level)))
	ligneFiche("")
	ligneFiche(fmt.Sprintf("  PV   %s   %3d / %d", pvBar(joueur.HP, joueur.MaxHP), joueur.HP, joueur.MaxHP))
	ligneFiche(fmt.Sprintf("  XP   %s   %3d / %d", pvBar(joueur.XP, joueur.XPToNextLevel()), joueur.XP, joueur.XPToNextLevel()))
	ligneFiche(deuxCotes(fmt.Sprintf("  Or   %d pièces", joueur.Gold), fmt.Sprintf("Sac  %2d / %d  ", joueur.Inventory.Size(), MaxInventorySize)))
	fmt.Println("╚" + bord + "╝")
	fmt.Println("  " + prochainObjectif(joueur))
}

// prochainObjectif indique la prochaine zone que le joueur peut débloquer.
func prochainObjectif(joueur *Player) string {
	for _, z := range allZones() {
		if z.niveauMin > joueur.Level {
			return fmt.Sprintf("→ Prochaine zone : %s (niveau %d)", z.nom, z.niveauMin)
		}
	}
	return "→ Toutes les zones sont débloquées !"
}

// ligneFiche affiche une ligne du cadre, complétée par des espaces.
func ligneFiche(texte string) {
	fmt.Printf("║%-*s║\n", largeurFiche, texte)
}

// centrer place le texte au milieu du cadre.
func centrer(texte string) string {
	marge := (largeurFiche - utf8.RuneCountInString(texte)) / 2
	if marge < 0 {
		marge = 0
	}
	return strings.Repeat(" ", marge) + texte
}

// deuxCotes place un texte à gauche et un autre à droite du cadre.
func deuxCotes(gauche, droite string) string {
	espace := largeurFiche - utf8.RuneCountInString(gauche) - utf8.RuneCountInString(droite)
	if espace < 1 {
		espace = 1
	}
	return gauche + strings.Repeat(" ", espace) + droite
}

// menuObjets permet d'utiliser les objets de l'inventaire hors combat.
func menuObjets(joueur *Player) {
	for {
		clearScreen()
		fmt.Println("╔══════════════════════════════╗")
		fmt.Println("║        UTILISER OBJET        ║")
		fmt.Println("╚══════════════════════════════╝")
		fmt.Printf("%s %d/%d PV\n", pvBar(joueur.HP, joueur.MaxHP), joueur.HP, joueur.MaxHP)
		fmt.Println()

		if joueur.Inventory.Size() == 0 {
			fmt.Println("Ton inventaire est vide...")
			attendreEntree()
			return
		}

		for i, objet := range joueur.Inventory.Items {
			fmt.Printf("[%d] %s - %s\n", i+1, objet.Name, objet.Description)
		}
		fmt.Println("[0] Retour")

		choix := lireChoix(0, joueur.Inventory.Size())
		if choix == 0 {
			return
		}

		fmt.Println()
		joueur.Inventory.UseItem(choix-1, joueur)
		attendreEntree()
	}
}

// reanimer ramène le joueur au village après une défaite,
// contre la moitié de son or.
func reanimer(joueur *Player) {
	perte := joueur.Gold / 2
	joueur.RemoveGold(perte)
	joueur.Heal(joueur.MaxHP / 2)

	fmt.Println()
	fmt.Println("💫 Le médecin du village t'a retrouvé et soigné...")
	fmt.Println("Tu perds", perte, "pièces d'or et reviens avec", joueur.HP, "PV.")
	time.Sleep(1 * time.Second)
}

// =========================
// SAISIE CLAVIER
// =========================

// lireLigne lit une ligne sur l'entrée standard, octet par octet.
// On n'utilise pas bufio ici : il lirait en avance et "volerait" la saisie
// des fmt.Scan utilisés dans le reste du jeu.
func lireLigne() string {
	var sb strings.Builder
	buf := make([]byte, 1)

	for {
		n, err := os.Stdin.Read(buf)
		if n == 0 || err != nil {
			// Entrée fermée (Ctrl+Z / Ctrl+D) : on quitte proprement.
			fmt.Println()
			fmt.Println("Au revoir !")
			os.Exit(0)
		}
		if buf[0] == '\n' {
			break
		}
		sb.WriteByte(buf[0])
	}

	return strings.TrimSpace(sb.String())
}

// lireChoix demande un nombre entre min et max jusqu'à obtenir une réponse valide.
func lireChoix(min, max int) int {
	for {
		fmt.Print("> ")
		ligne := lireLigne()

		// Ligne vide : reste d'un fmt.Scan précédent, on redemande sans message.
		if ligne == "" {
			continue
		}

		choix, err := strconv.Atoi(ligne)
		if err == nil && choix >= min && choix <= max {
			return choix
		}

		fmt.Printf("Choix invalide, entre un nombre entre %d et %d.\n", min, max)
	}
}

func attendreEntree() {
	fmt.Println()
	fmt.Print("Appuie sur Entrée pour continuer...")
	lireLigne()
}
