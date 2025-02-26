package main

import (
	"fmt"

	"github.com/mjh1207/gopokedex/internal/pokeapi"
)

func commandMap(conf *config) error {
	areas, err := pokeapi.GetLocationArea(conf.next)

	// Set new config state - Next
	if err != nil {
		return fmt.Errorf("GetLocationArea failed - Error: %v", err)
	}
	if areas.Next == nil {
		conf.next = ""
	} else {
		conf.next = *areas.Next
	}

	// Set new config state - Previous
	if areas.Previous == nil {
		conf.previous = ""
	} else {
		conf.previous = *areas.Previous
	}

	// Print all result locations
	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(conf *config) error {
	areas, err := pokeapi.GetLocationArea(conf.previous)
	// Set new config state - Next
	if err != nil {
		return fmt.Errorf("GetLocationArea failed - Error: %v", err)
	}
	if areas.Next == nil {
		conf.next = ""
	} else {
		conf.next = *areas.Next
	}

	// Set new config state - Previous
	if areas.Previous == nil {
		conf.previous = ""
	} else {
		conf.previous = *areas.Previous
	}

	// Print all result locations
	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}
	return nil
}