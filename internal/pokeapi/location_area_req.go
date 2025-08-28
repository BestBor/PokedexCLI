package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (Pokelocations, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokelocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokelocations{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return Pokelocations{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokelocations{}, err
	}

	locations := Pokelocations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Pokelocations{}, nil
	}
	return locations, nil

}
