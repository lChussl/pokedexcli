package main

import (
    "time"
    "github.com/lChussl/pokedexcli/internal/pokeapi"
)

func main() {
    pokeClient := pokeapi.NewClient(10 * time.Second)
    cfg := &config{
        client: pokeClient,
    }

    terminalReader(cfg)
}
