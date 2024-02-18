package domain

import (
	"context"
	UrlShortner "shortner"
)

const (
	EHashConflict string = "10000"
	ENotFound     string = "10001"
)

// KeyValueRepository is used to store shortened url and their mappings
type KeyValueRepository interface {
	Set(ctx context.Context, key string, value string) UrlShortner.Error
	Get(ctx context.Context, key string) (*Url, UrlShortner.Error)
	Delete(ctx context.Context, key string) UrlShortner.Error
}
