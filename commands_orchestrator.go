package main

import (
	"github.com/mjh1207/gopokedex/internal/pokeapi"
)

type cliCommand struct {
	name string
	description string
	callback func(conf *config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	next *string
	previous *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Get next 20 Pokemon area locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get previous 20 Pokemon area locations",
			callback: commandMapB,
		},
	}
}

