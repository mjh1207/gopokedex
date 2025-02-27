package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationArea(pageURL *string) (LocationArea, error){
	// Set request url
	url := baseUrl + "/location-area"
	if pageURL != nil { url = *pageURL }

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// Make GET request and verify response
	res, err := c.httpClient.Do(req)
	if err != nil || res.StatusCode > 299 {
		fmt.Println("FAILED TO GET DATA")
		return LocationArea{}, fmt.Errorf("unable to GET from %s - Status Code: %v", url, res.StatusCode)
	}
	defer res.Body.Close()

	//Decode response into location struct and return
	var areas LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return LocationArea{}, fmt.Errorf("unable to decode json - Error: %v", err)
	}

	return areas, nil
}