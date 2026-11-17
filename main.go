package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	id := "1"





routeur()











	//affiche les details du pokemon 
	character, err := getCharacter(client, id)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			log.Fatalf("Personnage %q introuvable, vérifiez l'id.", id)
		case errors.Is(err, ErrInvalidID):
			log.Fatalf("L'id %q n'est pas valide (un nombre est attendu).", id)
		case errors.Is(err, ErrRequestCreation):
			log.Fatalf("Problème interne lors de la création de la requête.")
		case errors.Is(err, ErrRequestSend):
			log.Fatalf("Impossible de joindre l'API, vérifiez votre connexion.")
		case errors.Is(err, ErrDecode):
			log.Fatalf("La réponse de l'API est illisible.")
		case errors.Is(err, ErrUnexpectedState):
			log.Fatalf("L'API a renvoyé une réponse inattendue.")
		default:
			log.Fatalf("Erreur inconnue : %v", err)
		}
	}

	fmt.Println(*character)





/*
//affiche la liste des pokemons
	list, err := getList(client)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			log.Fatalf("Liste introuvable.")
		case errors.Is(err, ErrRequestCreation):
			log.Fatalf("Problème interne lors de la création de la requête.")
		case errors.Is(err, ErrRequestSend):
			log.Fatalf("Impossible de joindre l'API, vérifiez votre connexion.")
		case errors.Is(err, ErrDecode):
			log.Fatalf("La réponse de l'API est illisible.")
		case errors.Is(err, ErrUnexpectedState):
			log.Fatalf("L'API a renvoyé une réponse inattendue.")
		default:
			log.Fatalf("Erreur inconnue : %v", err)
		}
	}

	fmt.Println(*list) */
}
