package crawling

import "luminnovel/internal/entity"

type resourceProvider interface {
	GetSheetValues(spreadsheetID string, range_ string) ([]entity.CrawlingItem, map[string]int, error)
	UpdateSheetValues(spreadsheetID string, range_ string, mapHeader map[string]int, payload []entity.CrawlingItem) (int64, error)
}

type service struct {
	resource resourceProvider
}

func NewService(rsc resourceProvider) *service {
	return &service{rsc}
}
