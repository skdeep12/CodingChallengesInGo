package repository

import (
	"context"
	"fmt"
	UrlShortner "shortner"
	"shortner/domain"
)

const (
	EHashConflict string = "10000"
	ENotFound     string = "10001"
)

type inMemoryKVRepository struct {
	m map[string]*domain.Url
}

func NewInMemoryStore() domain.KeyValueRepository {
	return &inMemoryKVRepository{
		m: make(map[string]*domain.Url),
	}
}

func (r *inMemoryKVRepository) Set(ctx context.Context, key string, value string) UrlShortner.Error {
	if val, ok := r.m[key]; !ok {
		r.m[key] = &domain.Url{
			Val:  value,
			Hits: 0,
		}
	} else if val.Val != value {
		return UrlShortner.NewError(EHashConflict, fmt.Sprintf("%s already hashed at %s", val.Val, key))
	}
	return nil
}

func (r *inMemoryKVRepository) Get(ctx context.Context, key string) (*domain.Url, UrlShortner.Error) {
	if val, ok := r.m[key]; !ok {
		return nil, UrlShortner.NewError(ENotFound, fmt.Sprintf("%s does not exist in the store", key))
	} else {
		return val, nil
	}
}
func (r *inMemoryKVRepository) Delete(ctx context.Context, key string) UrlShortner.Error {
	return nil
}
