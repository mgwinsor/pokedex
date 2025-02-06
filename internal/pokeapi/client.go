package pokeapi

import (
	"net/http"
	"time"

	"github.com/mgwinsor/pokedex/internal/pokecache"
)

type Client struct {
	pokeCache  pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		pokeCache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
