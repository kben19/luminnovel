package product

import (
	"context"

	"luminnovel/internal/entity"
)

type usecaseProductProvider interface {
	CrawlingAllProductSeries(ctx context.Context) error
	CrawlingProductSeries(ctx context.Context, title entity.ProductTitle, source entity.SiteSource) error
}

type Handler struct {
	usecaseProduct usecaseProductProvider
}

func New(productUC usecaseProductProvider) *Handler {
	return &Handler{productUC}
}
