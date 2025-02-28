package main

import "fmt"

func commandPokedex(conf *config, param string) error {
	if len(conf.pokedex) == 0 {
		fmt.Println("You have not caught any pokemon yet.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for pokemon := range conf.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}