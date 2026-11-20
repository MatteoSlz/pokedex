package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pokedex/models"
	"pokedex/routes"
	"pokedex/services"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	id := "1"

	routes.Routeur()

	character, err := services.GetCharacter(client, id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			log.Fatalf("Personnage %q introuvable, vérifiez l'id.", id)
		case errors.Is(err, models.ErrInvalidID):
			log.Fatalf("L'id %q n'est pas valide (un nombre est attendu).", id)
		case errors.Is(err, models.ErrRequestCreation):
			log.Fatalf("Problème interne lors de la création de la requête.")
		case errors.Is(err, models.ErrRequestSend):
			log.Fatalf("Impossible de joindre l'API, vérifiez votre connexion.")
		case errors.Is(err, models.ErrDecode):
			log.Fatalf("La réponse de l'API est illisible.")
		case errors.Is(err, models.ErrUnexpectedState):
			log.Fatalf("L'API a renvoyé une réponse inattendue.")
		default:
			log.Fatalf("Erreur inconnue : %v", err)
		}
	}

	fmt.Println(*character)
}
