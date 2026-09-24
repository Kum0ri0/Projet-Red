# ⚔️ Projet Red

Un petit jeu de rôle (RPG) **dans le terminal**, écrit en **Go**.
Tu crées ton personnage, tu récoltes des ressources au village, tu forges des armes et tu combats des monstres jusqu'au Dragon. 🐉

Projet Ynov réalisé par **Noah Bernadet**, **Jade Delahaye** et **Mathys Durand**.

---

## ▶️ Lancer le jeu

Il faut avoir [Go](https://go.dev/dl/) installé (version 1.27 ou plus).

```bash
git clone https://github.com/Kum0ri0/Projet-Red.git
cd Projet-Red
go run .
```

Pour jouer : tape le **numéro** du choix, puis appuie sur **Entrée**.

> Le jeu est fait pour **Windows**. Utilise **Windows Terminal** pour bien voir les emojis.

---

## 🎮 Comment se déroule une partie

```
Écran titre
   ↓
Création du personnage (nom + classe)
   ↓
Menu principal
   ├── [1] Village      → récolter, forger, acheter, se soigner, quêtes
   ├── [2] Aventure     → choisir une zone et combattre un monstre
   ├── [3] Personnage   → voir ses stats et son inventaire
   ├── [4] Objets       → utiliser une potion
   └── [5] Quitter
```

---

## 🧙 1. Le personnage

Au début, tu choisis ton **nom** et ta **classe** :

| Classe | Bonus |
|---|---|
| 🛡️ Guerrier | +20 PV max (120 PV) |
| 🔮 Mage | Connaît déjà **Boule de Feu** |
| 🗡️ Voleur | +30 pièces d'or |

Tout le monde commence avec :
- 100 PV, niveau 1
- le sort **Coup de poing** (8 dégâts)
- **2 potions de soin**
- **50 pièces d'or**

**Monter de niveau :** il faut `niveau × 100` XP. À chaque niveau, tu gagnes **+20 PV max** et tu es soigné complètement.

**Si tu meurs :** le médecin te récupère. Tu perds **la moitié de ton or** et tu reviens avec la moitié de tes PV.

---

## 🎒 2. L'inventaire

- Il peut contenir **10 objets maximum**.
- C'est le **même inventaire** partout : ce que tu récoltes au village sert en combat, et le butin des monstres peut être revendu au village.
- Tu peux le voir dans **Menu principal → [3] Personnage**.
- Tu peux utiliser tes potions dans **Menu principal → [4] Objets** ou **en combat → [2] Objets**.

---

## 🧪 3. Les objets

| Type | Objets | À quoi ça sert |
|---|---|---|
| 🧪 **Potions** | Potion de soin, Potion de poison | Soin : +50 PV. Poison : −30 PV |
| 🗡️ **Armes** | Épée en bois, Épée en métal, Hache, Épée, Dague, Bâton de mage, Épée du Dragon | Se fabriquent, se revendent ou se trouvent sur les monstres |
| 🪵 **Ressources** | Bois, Métal, Fourrure, Cuir, Peau | Servent à forger ou à revendre |
| 📖 **Grimoires** | Lame d'eau, Boule de Feu, Tempête, Blizzard, Lumière Divine, Souffle du Dragon | Apprennent un nouveau sort (à lire **en combat**) |

---

## 🏘️ 4. Le village de Marchang

| Lieu | Ce qu'on y fait |
|---|---|
| 🌿 **Prairies** | Récolter 3 bois |
| 🌲 **Forêt** | Récolter 3 bois |
| ⛏️ **Mine** | Récolter 2 métaux |
| 🔨 **Forgeron** | Épée en bois = **5 bois**. Épée en métal = **3 métal + 1 bois** |
| 💰 **Marchand** | Acheter et vendre (voir le tableau plus bas) |
| 🩺 **Médecin** | Remet tous tes PV au maximum |
| 📜 **Quêtes** | Voir ta progression et gagner de l'or |
| 🎒 **Inventaire** | Voir ce que tu portes |

### 💰 Prix du marchand

| Acheter | Prix | | Vendre | Prix |
|---|---|---|---|---|
| Bois | 10 or | | Bois | 5 or |
| Métal | 20 or | | Métal | 10 or |
| Grimoire : Lame d'eau | 40 or | | Épée en bois | 20 or |
| Grimoire : Boule de Feu | 80 or | | Épée en métal | 40 or |

### 📜 Quêtes

| Quête | Objectif | Récompense |
|---|---|---|
| Le bois du village | Récolter 10 bois | 50 or |
| Le métal de la mine | Récolter 10 métaux | 75 or |
| L'épée en bois | Fabriquer 1 épée en bois | 50 or |
| L'épée en métal | Fabriquer 1 épée en métal | 100 or |

---

## ⚔️ 5. Le combat

Le combat se joue **au tour par tour**. À ton tour, tu choisis :

| Choix | Effet |
|---|---|
| **[1] Attaque** | Tu lances un de tes sorts sur le monstre |
| **[2] Objets** | Tu bois une potion ou tu lis un grimoire |
| **[3] Fuir** | Tu quittes le combat |

Ensuite, **le monstre attaque** :
- il choisit une de ses attaques au hasard ;
- chaque attaque peut **rater** (selon sa précision) ;
- **tous les 3 tours**, il est **enragé 🔥** et fait **deux fois plus de dégâts**.

**Si tu gagnes :** tu reçois de l'**or**, de l'**XP** et parfois un **objet**.

### ✨ Les sorts

| Sort | Dégâts | Comment l'obtenir |
|---|:---:|---|
| Coup de poing | 8 | Au départ |
| Lame d'eau | 16 | Grimoire (marchand) |
| Boule de Feu | 18 | Classe Mage ou grimoire (marchand) |
| Tempête | 20 | Grimoire (Ogre) |
| Lumière Divine | 25 | Grimoire (Mage) |
| Blizzard | 30 | Grimoire (Mage) |
| Souffle du Dragon | 80 | Grimoire (Dragon) |

---

## 🗺️ 6. Les zones et les monstres

Certaines zones sont **bloquées 🔒** tant que tu n'as pas le niveau.

| Zone | Niveau | Monstres |
|---|:---:|---|
| 🌿 Prairie calme | 1 | Slime, Gobelin |
| 🌲 Forêt sombre | 3 | Sanglier, Loup |
| ⛰️ Montagne maudite | 5 | Mage, Ogre |
| 🌋 Antre du Dragon | 10 | Dragon |

| Monstre | PV | XP | Or | Butin possible |
|---|:---:|:---:|:---:|---|
| 🧊 Slime | 20 | 10 | 1–3 | — |
| 🧌 Gobelin | 50 | 20 | 4–7 | Dague (2 %) |
| 🐺 Loup | 70 | 30 | 6–8 | Fourrure (70 %) |
| 🐗 Sanglier | 90 | 20 | 10–13 | Cuir (40 %) |
| 🧙 Mage | 100 | 40 | 14–17 | Bâton de mage (40 %), Grimoire Blizzard (30 %), Grimoire Lumière Divine (10 %) |
| 👹 Ogre | 130 | 60 | 18–21 | Peau (30 %), Grimoire Tempête (40 %) |
| 🐉 Dragon | 200 | 100 | 30–35 | Épée du Dragon (5 %), Grimoire Souffle du Dragon (5 %) |

---

## 🗂️ 7. Les fichiers du projet

| Fichier | Rôle |
|---|---|
| `main.go` | Lance le jeu : écran titre, création du perso, menu principal, lecture du clavier |
| `player.go` | Le personnage : PV, niveau, XP, or, sorts |
| `items.go` | La liste de tous les objets et grimoires |
| `inventary.go` | L'inventaire : ajouter, retirer, afficher, utiliser un objet |
| `village.go` | Le menu du village |
| `prairies.go` · `foret.go` · `mine.go` | Récolte de ressources |
| `forgeron.go` | Fabrication des épées |
| `marchand.go` | Achat et vente |
| `medecin.go` | Soins |
| `quetes.go` | Les quêtes et leurs récompenses |
| `zone.go` | Les zones et le choix de la zone |
| `monster.go` | Les monstres, leurs attaques et leur butin |
| `combat.go` | Le déroulement d'un combat |

---

## 👥 L'équipe

| Membre | Partie |
|---|---|
| **Noah Bernadet** | Combat, monstres, zones |
| **Jade Delahaye** | Personnage, objets, inventaire |
| **Mathys Durand** | Village, lieux, forgeron, marchand, quêtes |
