
package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/api"
	"pokedex/models"
	"strings"
)

func GetCharacter(client *http.Client, id string) (*models.Character, error) {
	baseURL := strings.TrimRight(api.BaseURL, "/")
	endpoint := fmt.Sprintf("%s/%v", baseURL, id)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, models.ErrRequestCreation
	}
	request.Header.Set("User-Agent", "Pokedex-App/1.0")
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrRequestSend, err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, fmt.Errorf("%w (status %d)", models.ErrUnexpectedState, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrDecode, err)
	}

	var character models.Character
	if err := json.Unmarshal(body, &character); err != nil {
		preview := string(body)
		if len(preview) > 200 {
			preview = preview[:200]
		}
		return nil, fmt.Errorf("%w (URL: %s) : %v | Réponse reçue : %s", models.ErrDecode, endpoint, err, preview)
	}

	return &character, nil
}

func GetList(client *http.Client) (*models.List, error) {
	baseURL := strings.TrimRight(api.BaseURL, "/")
	endpoint := fmt.Sprintf("%s?limit=151&offset=0", baseURL)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, models.ErrRequestCreation
	}
	request.Header.Set("User-Agent", "Pokedex-App/1.0")
	request.Header.Set("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrRequestSend, err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, models.ErrNotFound
	default:
		return nil, fmt.Errorf("%w (status %d)", models.ErrUnexpectedState, response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", models.ErrDecode, err)
	}

	var list models.List
	if err := json.Unmarshal(body, &list); err != nil {
		preview := string(body)
		if len(preview) > 200 {
			preview = preview[:200]
		}
		return nil, fmt.Errorf("%w (URL: %s) : %v | Réponse reçue : %s", models.ErrDecode, endpoint, err, preview)
	}

	return &list, nil
}
