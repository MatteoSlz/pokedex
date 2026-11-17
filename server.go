package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

func routeur() {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		liste, err := getList(&client)
		if err != nil {
			http.Error(w, "Erreur lors de la récupération", 500)
			return
		}

		if len(liste.Results) > 151 {
			liste.Results = liste.Results[:151]
		}

		var pokemons []PokemonItem
		for _, p := range liste.Results {
			parts := strings.Split(strings.TrimSuffix(p.URL, "/"), "/")
			id := parts[len(parts)-1]

			pokemons = append(pokemons, PokemonItem{
				ID:    id,
				Name:  p.Name,
				Image: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/" + id + ".png",

			})
		}

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Erreur template", 500)
			return
		}

		tmpl.Execute(w, pokemons)
	})

	mux.HandleFunc("GET /pokemon/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		perso, err := getCharacter(&client, id)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				http.Error(w, "Pokemon introuvable", 404)
				return
			}
			if errors.Is(err, ErrInvalidID) {
				http.Error(w, "ID invalide", 400)
				return
			}
			http.Error(w, "Erreur serveur", 500)
			return
		}

		tmpl, err := template.ParseFiles("templates/detail.html")
		if err != nil {
			http.Error(w, "Erreur template", 500)
			return
		}

		tmpl.Execute(w, perso)
	})

	fmt.Println("Serveur lancé sur le port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
