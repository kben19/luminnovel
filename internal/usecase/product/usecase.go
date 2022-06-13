package product

import (
	"context"

	"luminnovel/internal/service/crawling"
	"luminnovel/internal/service/rightstufanime"
)

type crawlingSheetManager interface {
	UpdateCrawlingSheet(payload []crawling.CrawlingPayload, selection crawling.CrawlingConfig) error
}

type rightStufAnimeManager interface {
	FetchProductDataByTitle(ctx context.Context, title string) ([]rightstufanime.Product, error)
}

type Usecase struct {
	crawlingSheetSvc crawlingSheetManager
	rightStufSvc     rightStufAnimeManager
}

func New(crawling crawlingSheetManager, rightStuf rightStufAnimeManager) *Usecase {
	return &Usecase{crawling, rightStuf}
}
