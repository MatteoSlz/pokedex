package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokedex/api"
	"pokedex/models"
)

func GetCharacter(client *http.Client, id string) (*models.Character, error) {
	endpoint := fmt.Sprintf("%s/%v", api.BaseURL, id)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, models.ErrRequestCreation
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, models.ErrRequestSend
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, models.ErrNotFound
	case http.StatusInternalServerError:
		return nil, models.ErrInvalidID
	default:
		return nil, models.ErrUnexpectedState
	}

	var character models.Character
	if err := json.NewDecoder(response.Body).Decode(&character); err != nil {
		return nil, models.ErrDecode
	}

	return &character, nil
}

func GetList(client *http.Client) (*models.List, error) {
	endpoint := fmt.Sprintf("%s?limit=10000&offset=0", api.BaseURL)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, models.ErrRequestCreation
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, models.ErrRequestSend
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, models.ErrNotFound
	case http.StatusInternalServerError:
		return nil, models.ErrInvalidID
	default:
		return nil, models.ErrUnexpectedState
	}

	var list models.List
	if err := json.NewDecoder(response.Body).Decode(&list); err != nil {
		return nil, models.ErrDecode
	}

	return &list, nil
}
