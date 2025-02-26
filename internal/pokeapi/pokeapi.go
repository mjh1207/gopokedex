package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const defaultAreaUrl string = "https://pokeapi.co/api/v2/location-area/"

type LocationArea struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string) (LocationArea, error){
	if url == "" { url = defaultAreaUrl }
	res, err := http.Get(url)
	if err != nil || res.StatusCode > 299 {
		fmt.Println("FAILED TO GET DATA")
		return LocationArea{}, fmt.Errorf("unable to GET from %s - Status Code: %v", url, res.StatusCode)
	}
	defer res.Body.Close()

	var areas LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&areas)
	if err != nil {
		return LocationArea{}, fmt.Errorf("unable to decode json - Error: %v", err)
	}

	return areas, nil
}