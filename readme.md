# Pokédex

Projet web en Go utilisant la PokéAPI pour afficher les 151 premiers Pokémon et leurs statistiques.

## Prérequis

- Go installé sur la machine

## Lancement

Depuis la racine du projet, exécuter la commande suivante :

```bash
go run .
```

Le serveur démarre sur `http://localhost:8080`.

## Structure

- `main.go` : lancement du serveur
- `routes/` : gestion des routes HTTP
- `services/` : requêtes vers l'API
- `models/` : structures de données et erreurs
- `templates/` : fichiers HTML
- `assets/` : fichiers CSS
