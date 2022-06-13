package product

import (
	"context"
	"luminnovel/internal/service/bookdepository"

	"luminnovel/internal/service/crawling"
	"luminnovel/internal/service/rightstufanime"
)

type crawlingSheetManager interface {
	UpdateCrawlingSheet(payload []crawling.CrawlingPayload, selection crawling.CrawlingConfig) error
}

type rightStufAnimeManager interface {
	FetchProductDataByTitle(ctx context.Context, title string) ([]rightstufanime.Product, error)
}

type bookDepositoryManager interface {
	FetchProductDataByTitle(ctx context.Context, title string) ([]bookdepository.Product, error)
}

type Usecase struct {
	bookDepoSvc      bookDepositoryManager
	crawlingSheetSvc crawlingSheetManager
	rightStufSvc     rightStufAnimeManager
}

func New(crawling crawlingSheetManager, rightStuf rightStufAnimeManager, bookDepo bookDepositoryManager) *Usecase {
	return &Usecase{bookDepo, crawling, rightStuf}
}
