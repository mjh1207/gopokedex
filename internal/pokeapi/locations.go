package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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
	url := baseUrl + "/location-area/?offset=0&limit=20"
	if pageURL != nil { url = *pageURL }

	// If url is in the cache, use that data
	val, ok := c.cache.Get(url)
	if ok {
		fmt.Println("****************USING CACHED DATA**************************")
		var areas LocationArea
		if err := json.Unmarshal(val, &areas); err != nil {
			return LocationArea{}, err
		}
		return areas, nil
	}

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

	// Unmarshal response into location struct and return
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	
	var areas LocationArea
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)
	return areas, nil
}