package client

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) GetPokemonByName(
	ctx context.Context,
	pokemonName string,
) (Pokemon, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiUrl+"/api/v2/pokemon/"+pokemonName,
		nil,
	)
	if err != nil {
		return Pokemon{}, defaultInternalServerError()
	}
	req.Header.Add("Accept", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, defaultInternalServerError()
	}
	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, defaultInternalServerError()
	}
	return pokemon, nil
}

func defaultInternalServerError() PokemonFetchErr {
	return PokemonFetchErr{
		Message:    "non-200 status code from the API",
		StatusCode: http.StatusInternalServerError,
	}
}
