package main

import "fmt"

func commandInspect(conf *config, pokemonName string) error {
	// If poemon is in the pokedex, display it, otherwise display pokemon not caught message
	pokemonData, ok := conf.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemonData.Name)
	fmt.Printf("Height: %d\n", pokemonData.Height)
	fmt.Printf("Weight: %d\n", pokemonData.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonData.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemonData.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}
	return nil
}