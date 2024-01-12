package main

import (
    "time"
    "github.com/lChussl/pokedexcli/internal/pokeapi"
)

func main() {
    client := NewClient(10 * time.Second)
    terminalReader()
}
