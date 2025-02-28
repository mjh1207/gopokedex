package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(conf *config, param string) error {
	// Get Pokemon Data
	pokemonData, err := conf.pokeapiClient.GetPokemonData(param)
	if err != nil {
		return err
	}

	// Get capture rate for pokemon
	speciesData, err := conf.pokeapiClient.GetSpeciesData(param)
	if err != nil {
		return err
	}


	// Attempt to catch
	fmt.Printf("Throwing a Pokeball at %s...\n", param)
	chance := int(math.Round((float64(speciesData.CaptureRate) / 255.00 * 100)))
	attempt := rand.Intn(100)
	if caught := chance >= attempt; caught {
		fmt.Printf("%s was caught!\n", pokemonData.Name)
		conf.pokedex[pokemonData.Name] = pokemonData
	} else {
		fmt.Printf("%s escaped!\n", pokemonData.Name)
	}

	return nil
}