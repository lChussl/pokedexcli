package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/lChussl/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
    name            string
    description     string
    callback func() error
}

type config struct {
    client  pokeapi.Client
    nextUrl *string
    prevUrl *string
}

func terminalReader(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()

        text := cleanInput(scanner.Text())

        if len(text) == 0 {
            continue
        }

        commandName := text[0]

        command, ok := getCommands()[commandName]
        if ok {
            err := command.callback()
            if err != nil {
                fmt.Println(err)
            }

            continue
        } else {
            fmt.Println("Unknown command. Type 'help' for a list of commands.")

            continue
        }
    }
}

func cleanInput(input string) []string {
    output := strings.TrimSpace(input)
    words := strings.Fields(output)

    return words
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the pokedex",
            callback:    commandExit,
        },
        "map": {
            name:        "map",
            description: "Displays the first 20 areas of the map",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays the last 20 areas of the map",
            callback:    commandMapb,
        },
    }
}
