package crawling

import "luminnovel/internal/repository/googlesheet"

type googleSheetRepositoryProvider interface {
	Get(spreadsheetID string, range_ string) ([][]interface{}, error)
	Update(spreadsheetId string, range_ string, valuerange googlesheet.ValueRange) (int64, error)
}

type resource struct {
	repository googleSheetRepositoryProvider
}

func NewResource(repo googleSheetRepositoryProvider) *resource {
	return &resource{repo}
}
