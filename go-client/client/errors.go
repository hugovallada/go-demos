package client

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrFetchingPokemon = errors.New("failed to fetch pokemon")
)

type PokemonFetchErr struct {
	StatusCode int
	Message    string
}

func (e PokemonFetchErr) Error() string {
	return fmt.Sprintf("failed to fetch pokemon: %s with status code: %d", e.Message, http.StatusInternalServerError)
}
