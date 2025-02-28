package main

import (
	"time"

	"github.com/mjh1207/gopokedex/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		pokeapiClient: pokeClient,
		pokedex: map[string]pokeapi.Pokemon{},
	}
	startRepl(conf)
}