package pokeapi

import (
	"net/http"
	"time"

	"github.com/Bayan2019/pokedexcli/internal/pokecache"
)

// You'll need to use the PokeAPI location-area endpoint to get the location areas.
const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
