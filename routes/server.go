
package routes

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"pokedex/models"
	"pokedex/services"
	"strings"
	"time"
)

func Routeur() {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	mux := http.NewServeMux()

	// FICHIERS CSS
	mux.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("assets"))))

	// PAGE D'ACCUEIL
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		liste, err := services.GetList(&client)
		if err != nil {
			log.Printf("ERREUR GetList : %v", err)
			http.Error(w, "Erreur lors de la récupération : "+err.Error(), 500)
			return
		}

		if len(liste.Results) > 151 {
			liste.Results = liste.Results[:151]
		}

		var pokemons []models.PokemonItem

		for _, p := range liste.Results {
			parts := strings.Split(strings.TrimSuffix(p.URL, "/"), "/")
			id := parts[len(parts)-1]

			pokemons = append(pokemons, models.PokemonItem{
				ID:    id,
				Name:  p.Name,
				Image: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/" + id + ".png",
			})
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("ERREUR Template index : %v", err)
			http.Error(w, "Erreur template : "+err.Error(), 500)
			return
		}

		if err := tmpl.Execute(w, pokemons); err != nil {
			log.Printf("ERREUR Execute index : %v", err)
		}
	})

	// PAGE DETAIL POKEMON
	mux.HandleFunc("/pokemon/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/pokemon/")

		if id == "" || strings.Contains(id, "/") {
			http.NotFound(w, r)
			return
		}

		perso, err := services.GetCharacter(&client, id)
		if err != nil {
			log.Printf("ERREUR GetCharacter : %v", err)

			if errors.Is(err, models.ErrNotFound) {
				http.Error(w, "Pokemon introuvable", 404)
				return
			}

			if errors.Is(err, models.ErrInvalidID) {
				http.Error(w, "ID invalide", 400)
				return
			}

			http.Error(w, "Erreur serveur : "+err.Error(), 500)
			return
		}

		tmpl, err := template.ParseFiles("templates/detail.html")
		if err != nil {
			log.Printf("ERREUR Template detail : %v", err)
			http.Error(w, "Erreur template : "+err.Error(), 500)
			return
		}

		if err := tmpl.Execute(w, perso); err != nil {
			log.Printf("ERREUR Execute detail : %v", err)
		}
	})

	fmt.Println("Serveur lancé sur le port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
