package main 

import (
	"errors"
)

const baseURL = "https://pokeapi.co/api/v2/pokemon"

var (
	ErrRequestCreation = errors.New("impossible de créer la requête")
	ErrRequestSend     = errors.New("impossible d'envoyer la requête")
	ErrNotFound        = errors.New("ressource introuvable")
	ErrInvalidID       = errors.New("identifiant non valide")
	ErrUnexpectedState = errors.New("réponse inattendue de l'API")
	ErrDecode          = errors.New("impossible de décoder la réponse")
)
