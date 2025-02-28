package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetBatchLocations(pageURL *string) (RespBatchLocations, error){
	// Set request url
	url := baseUrl + locationAPI + "/?offset=0&limit=20"
	if pageURL != nil { url = *pageURL }

	// If url is in the cache, use that data
	val, ok := c.cache.Get(url)
	if ok {
		var areas RespBatchLocations
		if err := json.Unmarshal(val, &areas); err != nil {
			return RespBatchLocations{}, err
		}
		return areas, nil
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespBatchLocations{}, err
	}

	// Send GET request and verify response
	res, err := c.httpClient.Do(req)
	if err != nil || res.StatusCode > 299 {
		return RespBatchLocations{}, err
	}
	defer res.Body.Close()

	// Unmarshal response into batch location struct and return
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespBatchLocations{}, err
	}
	
	var areas RespBatchLocations
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return RespBatchLocations{}, err
	}

	c.cache.Add(url, data)
	return areas, nil
}

func (c *Client) GetLocationData(areaName string) (RespLocationData, error) {
	url := fmt.Sprintf("%s%s/%s",baseUrl, locationAPI, areaName)

	// if url is in cache, use that data
	val, ok := c.cache.Get(url)
	if ok {
		var locationData RespLocationData
		if err := json.Unmarshal(val, &locationData); err != nil {
			return RespLocationData{}, err
		}
		return locationData, nil
	}

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationData{}, err
	}

	// Send GET request and verify response 
	res, err := c.httpClient.Do(req)
	if err != nil || res.StatusCode > 299 {
		return RespLocationData{}, err
	}
	defer res.Body.Close()
	
	// Unmarshal response into location data struct and return
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationData{}, err
	}

	var areaInfo RespLocationData
	err = json.Unmarshal(data, &areaInfo)
	if err != nil {
		return RespLocationData{}, err
	}

	c.cache.Add(url, data)
	return areaInfo, nil
}