package rightstufanime

import (
	"context"

	"luminnovel/internal/entity"
)

type resourceProvider interface {
	FindSeriesByTitleFromDB(ctx context.Context, collectionName string, title string) ([]entity.Source, error)
	GetRequestFromHTTP(ctx context.Context, params entity.Source) (entity.ProductRightStufAnime, error)
}

type service struct {
	resource resourceProvider
}

func NewService(rsc resourceProvider) *service {
	return &service{rsc}
}
