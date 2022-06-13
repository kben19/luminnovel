package googlesheet

import (
	"errors"
	"google.golang.org/api/sheets/v4"
)

type repository struct {
	service *sheets.Service
}

func New(srv *sheets.Service) *repository {
	return &repository{srv}
}

func (repo *repository) Get(spreadsheetID string, range_ string) ([][]interface{}, error) {
	resp, err := repo.service.Spreadsheets.Values.Get(spreadsheetID, range_).Do()
	if err != nil {
		return nil, err
	}
	if len(resp.Values) == 0 {
		return nil, errors.New("no data found")
	}

	result := make([][]interface{}, len(resp.Values))
	for index, row := range resp.Values {
		result[index] = make([]interface{}, len(row))
		for indexRow, value := range row {
			result[index][indexRow] = value
		}
	}
	return result, nil
}

func (repo *repository) Update(spreadsheetId string, range_ string, valuerange ValueRange) (int64, error) {
	resp, err := repo.service.Spreadsheets.Values.Update(spreadsheetId, range_, &sheets.ValueRange{
		MajorDimension: string(valuerange.MajorDimension),
		Range:          valuerange.Range,
		Values:         valuerange.Values,
	}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return 0, err
	}
	return resp.UpdatedCells, nil
}
