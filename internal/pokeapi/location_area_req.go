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

	data, ok := c.cache.Get(fullURL)
	if ok {
		locations := Pokelocations{}
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return Pokelocations{}, nil
		}
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

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokelocations{}, err
	}

	locations := Pokelocations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Pokelocations{}, nil
	}

	c.cache.Add(fullURL, data)

	return locations, nil

}

func (c *Client) GetLocationArea(locationName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationName
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		locations := Pokelocations{}
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return LocationArea{}, nil
		}
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locations := LocationArea{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationArea{}, nil
	}

	c.cache.Add(fullURL, data)

	return locations, nil
}
