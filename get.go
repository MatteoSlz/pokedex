package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getCharacter(client *http.Client, id string) (*Character, error) {
	endpoint := fmt.Sprintf("%s/%v", baseURL, id)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, ErrRequestCreation
	}


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

	var character Character
	if err := json.NewDecoder(response.Body).Decode(&character); err != nil {
		return nil, ErrDecode
	}

	return &character, nil
}

func getList(client *http.Client) (*List, error) {
	endpoint := fmt.Sprintf("%s?limit=10000&offset=0", baseURL)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, ErrRequestCreation
	}

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

