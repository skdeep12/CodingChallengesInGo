package domain

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	UrlShortner "shortner"
	"time"
)

// Shortener provides the contract to shorten a given long url
type Shortener interface {
	Shorten(ctx context.Context, url string) (string, UrlShortner.Error)
}

// Resolver provides the contract to resolve a shortened url
type Resolver interface {
	// Resolve returns a long url for corresponding short url, if it exists in the system.
	// it throws ENotFound error
	Resolve(ctx context.Context, shortUrl string) (string, UrlShortner.Error)
}

type ShortenerResolver interface {
	Shortener
	Resolver
}

type Url struct {
	Val    string
	Hits   int
	Expiry *time.Time
}

func (u *Url) IncreaseHits() {
	u.Hits++
}

// testShortener implements ShortenerResolver interface
type testShortener struct {
	store KeyValueRepository
}

func NewTestShortener(store KeyValueRepository) ShortenerResolver {
	return &testShortener{store: store}
}

type depth struct{}

func (t *testShortener) Shorten(ctx context.Context, url string) (string, UrlShortner.Error) {
	hash := t.getHash(ctx, url)
	d := ctx.Value(depth{})
	if d != nil && d.(int) > 4 {
		return "", UrlShortner.NewError(EHashConflictResolutionDepthExceeded, fmt.Sprintf("depth is %d", d.(int)))
	}
	if err := t.store.Set(ctx, hash, url); err != nil {
		if err.Code() == EHashConflict {
			if d == nil {
				ctx = context.WithValue(ctx, depth{}, 1)
			} else {
				ctx = context.WithValue(ctx, depth{}, d.(int)+1)
			}
			hash, err = t.Shorten(ctx, url+time.Now().String())
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return hash, nil
}

func (t *testShortener) getHash(ctx context.Context, url string) string {
	h := sha256.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(bs)[:6]
}

func (t *testShortener) Resolve(ctx context.Context, shortUrl string) (string, UrlShortner.Error) {
	if resolvedUrl, err := t.store.Get(ctx, shortUrl); err != nil {
		return "", err
	} else {
		resolvedUrl.IncreaseHits()
		return resolvedUrl.Val, nil
	}
}
