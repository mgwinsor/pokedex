package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mgwinsor/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient   *pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		if cmd, ok := getCommands()[commandName]; ok {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Shows the help menu",
			callback:    commandHelp,
		},
		"catch": {
			name:        "explore",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location on the map",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the details of a pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all of your caught pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}
