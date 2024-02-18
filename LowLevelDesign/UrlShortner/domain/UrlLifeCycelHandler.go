package domain

import (
	"context"
	UrlShortner "shortner"
)

type UrlLifeCycleHandler interface {
	Delete(ctx context.Context, shortUrl string) UrlShortner.Error
}

// urlLifeCycleHandlerImpl implements UrlLifeCycleHandler interface
type urlLifeCycleHandlerImpl struct {
	store KeyValueRepository
}

func NewUrlLifeCycleHandler(store KeyValueRepository) UrlLifeCycleHandler {
	return &urlLifeCycleHandlerImpl{
		store: store,
	}
}

func (u *urlLifeCycleHandlerImpl) Delete(ctx context.Context, shortUrl string) UrlShortner.Error {
	return u.store.Delete(ctx, shortUrl)
}
