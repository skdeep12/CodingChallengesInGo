package domain

import (
	"context"
	"github.com/stretchr/testify/assert"
	UrlShortner "shortner"
	"testing"
)

type KeyValueRepositoryTestImpl struct {
}

func (t *KeyValueRepositoryTestImpl) Set(ctx context.Context, key string, value string) UrlShortner.Error {
	return nil
}
func (t *KeyValueRepositoryTestImpl) Get(ctx context.Context, key string) (*Url, UrlShortner.Error) {
	return nil, UrlShortner.NewError(ENotFound, "")
}
func (t *KeyValueRepositoryTestImpl) Delete(ctx context.Context, key string) UrlShortner.Error {
	return nil
}

// Test cases for url shortener
// when url doesn't exist send ENotFound
// when url exists, increment hit and get expanded url
func TestTestShortener_Resolve(t *testing.T) {
	s := NewTestShortener(&KeyValueRepositoryTestImpl{})
	_, err := s.Resolve(context.Background(), "txt")
	assert.NotEqual(t, nil, err, "error should not be nil if short url does not exist")
	assert.Equal(t, ENotFound, err.Code(), "error should not be nil if short url does not exist")
}
