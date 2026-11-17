package main 

import (
	"errors"
)

const baseURL = "https://pokeapi.co/api/v2/pokemon"

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type TypeInfo `json:"type"`
}

type StatInfo struct {
	Name string `json:"name"`
}

type PokemonStat struct {
	BaseStat int      `json:"base_stat"`
	Stat     StatInfo `json:"stat"`
}

type Character struct {
	Id      int           `json:"id"`
	Name    string        `json:"name"`
	Sprites Sprites       `json:"sprites"`
	Types   []PokemonType `json:"types"`
	Stats   []PokemonStat `json:"stats"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type List struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}

var (
	ErrRequestCreation = errors.New("impossible de créer la requête")
	ErrRequestSend     = errors.New("impossible d'envoyer la requête")
	ErrNotFound        = errors.New("ressource introuvable")
	ErrInvalidID       = errors.New("identifiant non valide")
	ErrUnexpectedState = errors.New("réponse inattendue de l'API")
	ErrDecode          = errors.New("impossible de décoder la réponse")
)