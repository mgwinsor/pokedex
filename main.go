package main

import (
	"time"

	"github.com/mgwinsor/pokedex/internal/pokeapi"
)

const (
	timeout       = 5 * time.Second
	cacheInterval = 5 * time.Minute
)

func main() {
	pokeClient := pokeapi.NewClient(timeout, cacheInterval)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
