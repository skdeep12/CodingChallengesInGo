package domain

import (
	"context"
	UrlShortner "shortner"
)

type KeyValueRepository interface {
	Set(ctx context.Context, key string, value string) UrlShortner.Error
	Get(ctx context.Context, key string) (*Url, UrlShortner.Error)
	Delete(ctx context.Context, key string) UrlShortner.Error
}
