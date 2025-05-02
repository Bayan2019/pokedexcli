package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 4. PokeAPI
func (c *Client) ListLocationAreas(pageURL *string) (RespShallowLocations, error) {
	// You'll need to use the PokeAPI location-area endpoint to get the location areas.
	// Calling the endpoint without an id will return a batch of location areas.
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")

		locationAreasResp := RespShallowLocations{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationAreasResp, nil
	}

	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return RespShallowLocations{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationAreasResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")

		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	fmt.Println("Cache miss")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationArea, nil
}
