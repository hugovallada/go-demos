package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	testUrl = "my-test-url"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("happy path - can hit the api and return a pokemon", func(*testing.T) {
		myClient := NewClient()
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})

	t.Run("sad path - return an error when the pokemon does not exist", func(*testing.T) {
		myClient := NewClient()
		_, err := myClient.GetPokemonByName(context.Background(), "non-existant")
		assert.Error(t, err)
		assert.Equal(t, PokemonFetchErr{Message: "non-200 status code from the API", StatusCode: http.StatusInternalServerError}, err)
	})

	t.Run("happy path - testing the WithAPIURL option func", func(*testing.T) {
		myClient := NewClient(
			WithAPIURL(testUrl),
		)

		assert.Equal(t, testUrl, myClient.apiUrl)
	})

	t.Run("happy path - testing the WithHTTPClient option func", func(*testing.T) {
		myClient := NewClient(
			WithAPIURL(testUrl),
			WithHTTPClient(
				&http.Client{
					Timeout: 1 * time.Second,
				},
			),
		)

		assert.Equal(t, testUrl, myClient.apiUrl)
		assert.Equal(t, 1*time.Second, myClient.httpClient.Timeout)
	})

	t.Run("happy path - able to hit locally running test server", func(*testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"name": "pikachu", "height":10}`)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, 10, poke.Height)
	})

	t.Run("sad path - able to handle 500 status from the API", func(*testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}),
		)
		defer ts.Close()

		myClient := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.Error(t, err)
		assert.Equal(t, 0, poke.Height)
	})
}
