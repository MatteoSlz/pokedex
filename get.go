package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getCharacter(client *http.Client, id string) (*Character, error) {
	endpoint := fmt.Sprintf("%s/%v", baseURL, id)

	// Préparation de la requête GET.
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, ErrRequestCreation
	}

	// Envoi de la requête.
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrRequestSend
	}
	defer response.Body.Close()

	// Vérification du code de statut HTTP.
	switch response.StatusCode {
	case http.StatusOK:
		// Succès : on continue vers le décodage.
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusInternalServerError:
		return nil, ErrInvalidID
	default:
		return nil, ErrUnexpectedState
	}

	// Décodage du JSON dans la structure Character.
	var character Character
	if err := json.NewDecoder(response.Body).Decode(&character); err != nil {
		return nil, ErrDecode
	}

	return &character, nil
}

func getList(client *http.Client) (*List, error) {
	// Ajout de ?limit=10000 pour récupérer tous les Pokémon
	endpoint := fmt.Sprintf("%s?limit=10000&offset=0", baseURL)

	// Préparation de la requête GET avec endpoint au lieu de baseURL
	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, ErrRequestCreation
	}

	// Le reste de votre fonction reste identique...
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrRequestSend
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusInternalServerError:
		return nil, ErrInvalidID
	default:
		return nil, ErrUnexpectedState
	}

	var list List
	if err := json.NewDecoder(response.Body).Decode(&list); err != nil {
		return nil, ErrDecode
	}

	return &list, nil
}

