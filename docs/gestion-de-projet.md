# Gestion de projet — Projet RED

> Si les encadrants fournissent leur propre document de gestion de projet, placez-le dans ce dossier `docs/` et complétez-le.

## Équipe et répartition

| Membre | Partie | Fichiers |
|---|---|---|
| Noah Bernadet | Combat, monstres, zones | `combat.go`, `monster.go`, `zone.go` |
| Jade Delahaye | Personnage, objets, inventaire | `player.go`, `items.go`, `inventary.go` |
| Mathys Durand | Village, lieux, économie, quêtes | `village.go`, `prairies.go`, `foret.go`, `mine.go`, `forgeron.go`, `marchand.go`, `medecin.go`, `quetes.go` |

`main.go` (menu principal, création du personnage) fait le lien entre les trois parties.

## Outils

- **Langage :** Go
- **Versionnement :** Git + GitHub (`https://github.com/Kum0ri0/Projet-Red`)
- **Communication :** [à compléter]

## Planning

| Période | Étape |
|---|---|
| 18/09/2026 | Début du projet, premiers commits |
| [à compléter] | Personnage et inventaire |
| [à compléter] | Village, marchand, forgeron |
| [à compléter] | Combat, monstres, zones |
| [à compléter] | Fusion des parties (un seul `Player`, un seul inventaire) |
| [à compléter] | Oral |

## Suivi des tâches du sujet

| Partie | Tâches | État |
|---|---|---|
| 1. Personnage | 1 à 12 | Faites (9 et 11 adaptées) |
| 2. Économie | 13 à 18 | 13 à 15 adaptées, 16 à 18 à faire |
| 3. Combat | 19 à 22 | Faites |
| Missions | 1 à 6 | 2, 3, 5 faites ; 1, 4, 6 à faire |

## Difficultés rencontrées

- **Travail à plusieurs sur les mêmes fichiers :** une modification de `player.go` a retiré l'inventaire du joueur et cassé la compilation du combat.
- **Deux inventaires en double :** le village utilisait ses propres variables (bois, argent, vie), séparées du `Player`. Ils ont été fusionnés.
- **Saisie clavier :** `fmt.Scan` réaffichait le menu en boucle sur une saisie invalide ; remplacé par `lireChoix()`