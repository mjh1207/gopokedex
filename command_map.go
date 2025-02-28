package main

import (
	"errors"
	"fmt"
)

func commandMap(conf *config, param string) error {
	areas, err := conf.pokeapiClient.GetBatchLocations(conf.next)
	if err != nil {
		return err
	}
	
	// Set new config state - Next
	conf.next = areas.Next
	conf.previous = areas.Previous

	// Print all result locations
	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(conf *config, param string) error {
	if conf.previous == nil {
		return errors.New("you are on the first page")
	}
	areas, err := conf.pokeapiClient.GetBatchLocations(conf.previous)
	if err != nil {
		return err
	}
	
	// Set new config state - Next
	conf.next = areas.Next
	conf.previous = areas.Previous

	// Print all result locations
	for _, result := range areas.Results {
		fmt.Println(result.Name)
	}
	return nil
}