package main

import "fmt"

func commandExplore(conf *config, param string) error {
	fmt.Printf("Exploring %s...\n", param)
	areaInfo, err := conf.pokeapiClient.GetLocationData(param)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range areaInfo.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}