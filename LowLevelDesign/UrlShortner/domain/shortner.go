package domain

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	UrlShortner "shortner"
	"time"
)

type Shortener interface {
	Shorten(ctx context.Context, url string) (string, UrlShortner.Error)
}

type Resolver interface {
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

type testShortener struct {
	store KeyValueRepository
}

func NewTestShortener(store KeyValueRepository) ShortenerResolver {
	return &testShortener{store: store}
}

func (t *testShortener) Shorten(ctx context.Context, url string) (string, UrlShortner.Error) {
	h := sha256.New()
	h.Write([]byte(url))
	bs := h.Sum(nil)
	fmt.Println(bs)
	encodedString := base64.StdEncoding.EncodeToString(bs)[:6]
	if err := t.store.Set(ctx, encodedString, url); err != nil {
		fmt.Println(err)
		return "", err
	}
	return encodedString, nil
}

func (t *testShortener) Resolve(ctx context.Context, shortUrl string) (string, UrlShortner.Error) {
	if resolvedUrl, err := t.store.Get(ctx, shortUrl); err != nil {
		return "", err
	} else {
		resolvedUrl.IncreaseHits()
		return resolvedUrl.Val, nil
	}
}
