package crawling

import (
	"luminnovel/internal/entity"
	"luminnovel/internal/repository/googlesheet"
	"strconv"
)

const (
	titleHeader                    = "Title"
	amazonHeader                   = "Amazon"
	rightStufHeader                = "RightStufAnime"
	inStockHeader                  = "InStockTrades"
	bookDepoHeader                 = "BookDepository"
	outOfStockRightStufHeader      = "Out of stock"
	outOfStockBookDepositoryHeader = "OOS BD"
	weightHeader                   = "Weight"

	totalColumn = 8
)

func (rsc *resource) GetSheetValues(spreadsheetID string, range_ string) ([]entity.CrawlingItem, map[string]int, error) {
	values, err := rsc.repository.Get(spreadsheetID, range_)
	if err != nil {
		return nil, nil, err
	}

	mapHeaderColumn := map[string]int{}
	header := values[0]
	for index, column := range header {
		switch column {
		case titleHeader:
			mapHeaderColumn[titleHeader] = index
		case amazonHeader:
			mapHeaderColumn[amazonHeader] = index
		case rightStufHeader:
			mapHeaderColumn[rightStufHeader] = index
		case inStockHeader:
			mapHeaderColumn[inStockHeader] = index
		case bookDepoHeader:
			mapHeaderColumn[bookDepoHeader] = index
		case outOfStockRightStufHeader:
			mapHeaderColumn[outOfStockRightStufHeader] = index
		case outOfStockBookDepositoryHeader:
			mapHeaderColumn[outOfStockBookDepositoryHeader] = index
		case weightHeader:
			mapHeaderColumn[weightHeader] = index
		}
	}

	if len(values) <= 1 {
		return nil, mapHeaderColumn, nil
	}
	values = values[1:]

	crawlings := make([]entity.CrawlingItem, len(values))
	for index, row := range values {
		rightStufStock, err := strconv.ParseBool(row[mapHeaderColumn[outOfStockRightStufHeader]].(string))
		if err != nil {
			return nil, nil, err
		}
		bookDepoStock, err := strconv.ParseBool(row[mapHeaderColumn[outOfStockBookDepositoryHeader]].(string))
		if err != nil {
			return nil, nil, err
		}
		weight, err := strconv.ParseFloat(row[mapHeaderColumn[weightHeader]].(string), 64)
		if err != nil {
			return nil, nil, err
		}
		crawlings[index] = entity.CrawlingItem{
			Title: row[mapHeaderColumn[titleHeader]].(string),
			Price: entity.CrawlingPrice{
				Amazon:         row[mapHeaderColumn[amazonHeader]].(string),
				RightStufAnime: row[mapHeaderColumn[rightStufHeader]].(string),
				InStockTrades:  row[mapHeaderColumn[inStockHeader]].(string),
				BookDepository: row[mapHeaderColumn[bookDepoHeader]].(string),
			},
			Stock: entity.CrawlingStock{
				Amazon:         false,
				RightStufAnime: rightStufStock,
				InStockTrades:  false,
				BookDepository: bookDepoStock,
			},
			Weight:       weight,
			ActualWeight: 0,
		}
	}
	return crawlings, mapHeaderColumn, nil
}

func (rsc *resource) UpdateSheetValues(spreadsheetID string, range_ string, mapHeader map[string]int, payload []entity.CrawlingItem) (int64, error) {
	valueRange := convertCrawlingItemPayload(payload, range_, mapHeader)

	updatedCells, err := rsc.repository.Update(spreadsheetID, range_, valueRange)
	if err != nil {
		return 0, err
	}
	return updatedCells, nil
}

func convertCrawlingItemPayload(payload []entity.CrawlingItem, range_ string, mapHeader map[string]int) googlesheet.ValueRange {
	values := make([][]interface{}, len(payload))
	for index, row := range payload {
		rowValue := make([]interface{}, totalColumn)
		rowValue[mapHeader[titleHeader]] = row.Title
		rowValue[mapHeader[amazonHeader]] = row.Price.Amazon
		rowValue[mapHeader[rightStufHeader]] = row.Price.RightStufAnime
		rowValue[mapHeader[inStockHeader]] = row.Price.InStockTrades
		rowValue[mapHeader[bookDepoHeader]] = row.Price.BookDepository
		rowValue[mapHeader[outOfStockRightStufHeader]] = row.Stock.RightStufAnime
		rowValue[mapHeader[outOfStockBookDepositoryHeader]] = row.Stock.BookDepository
		rowValue[mapHeader[weightHeader]] = row.Weight
		values[index] = rowValue
	}
	return googlesheet.ValueRange{
		MajorDimension: googlesheet.MajorDimensionRow,
		Range:          range_,
		Values:         values,
	}
}
