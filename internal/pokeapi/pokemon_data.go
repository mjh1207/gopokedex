package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetSpeciesData(pokeName string) (RespCaptureRate, error){
	// Set request url
	url := fmt.Sprintf("%s%s/%s", baseUrl, speciesAPI, pokeName)

	// if url is in cache, use that data
	val, ok := c.cache.Get(url)
	if ok {
		var data RespCaptureRate
		if err := json.Unmarshal(val, &data); err != nil {
			return RespCaptureRate{}, err
		}
		return data, nil
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespCaptureRate{}, err
	}

	// Send request and verify response
	res, err := c.httpClient.Do(req)
	if err != nil || res.StatusCode > 299 {
		return RespCaptureRate{}, fmt.Errorf("unable to get data for %s", pokeName)
	}
	defer res.Body.Close()

	// Unmarshal response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespCaptureRate{}, err
	}

	var pokemonInfo RespCaptureRate
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return RespCaptureRate{}, err
	}

	c.cache.Add(url, data)
	return pokemonInfo, nil
}

func (c *Client) GetPokemonData(pokeName string) (Pokemon, error){
	// Set request url
	url := fmt.Sprintf("%s%s/%s", baseUrl, pokemonAPI, pokeName)

	// if url is in cache, use that data
	val, ok := c.cache.Get(url)
	if ok {
		var data Pokemon
		if err := json.Unmarshal(val, &data); err != nil {
			return Pokemon{}, err
		}
		return data, nil
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// Send request and verify response
	res, err := c.httpClient.Do(req)
	if err != nil || res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("unable to get data for %s", pokeName)
	}
	defer res.Body.Close()

	// Unmarshal response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemonInfo Pokemon
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonInfo, nil
}