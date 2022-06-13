package crawling

import (
	"errors"
	"fmt"
	"luminnovel/internal/entity"
	"strconv"
)

const (
	spreadsheetIDCrawling = "19Zr7zFocTputKxo7-GJSmoueKUHghqkhEL7PxYtTw6c"
)

func (svc *service) UpdateCrawlingSheet(payload []CrawlingPayload, selection CrawlingConfig) error {
	sheetName := string(selection.Series)
	if sheetName == "" {
		return errors.New("invalid title name")
	}
	rangeGet := sheetName + "!A1:G"
	rangeSet := sheetName + "!A2:G"

	sheetValues, mapHeader, err := svc.resource.GetSheetValues(spreadsheetIDCrawling, rangeGet)
	if err != nil {
		return err
	}

	changes, priceChanges, stockChanges := updateSheetWithPayload(sheetValues, payload, selection)

	_, err = svc.resource.UpdateSheetValues(spreadsheetIDCrawling, rangeSet, mapHeader, sheetValues)
	if err != nil {
		return err
	}
	if changes > 0 {
		fmt.Printf("Updated %d cells in Sheet %s for source %s\n", changes, sheetName, selection.getSourcename())
		fmt.Printf("Price updated in volumes %s\n", arrayToString(priceChanges))
		fmt.Printf("Stock updated in volumes %s\n", arrayToString(stockChanges))
	} else {
		fmt.Println("No Updates")
	}
	return nil
}

func updateSheetWithPayload(sheetValues []entity.CrawlingItem, payload []CrawlingPayload, selection CrawlingConfig) (int, []int, []int) {
	var (
		countUpdate int
		stockUpdate []int
		priceUpdate []int
	)
	for _, item := range payload {
		var (
			priceChanged bool
			stockChanged bool
		)
		volumeIndex := item.Volume - 1

		if selection.Source.RightStufAnime {
			priceChanged = item.Price != sheetValues[volumeIndex].Price.RightStufAnime
			stockChanged = !item.InStock != sheetValues[volumeIndex].Stock.RightStufAnime
			sheetValues[volumeIndex].Price.RightStufAnime = item.Price
			sheetValues[volumeIndex].Stock.RightStufAnime = !item.InStock // row is out of stock format
		} else if selection.Source.InStockTrades {
			priceChanged = item.Price != sheetValues[volumeIndex].Price.InStockTrades
			stockChanged = !item.InStock != sheetValues[volumeIndex].Stock.InStockTrades
			sheetValues[volumeIndex].Price.InStockTrades = item.Price
			sheetValues[volumeIndex].Stock.InStockTrades = !item.InStock
		} else if selection.Source.Amazon {
			priceChanged = item.Price != sheetValues[volumeIndex].Price.Amazon
			stockChanged = !item.InStock != sheetValues[volumeIndex].Stock.Amazon
			sheetValues[volumeIndex].Price.Amazon = item.Price
			sheetValues[volumeIndex].Stock.Amazon = !item.InStock
			sheetValues[volumeIndex].Weight = item.Weight
		}

		// count updated rows
		if priceChanged || stockChanged {
			countUpdate++
		}
		if stockChanged {
			stockUpdate = append(stockUpdate, item.Volume)
		}
		if priceChanged {
			priceUpdate = append(priceUpdate, item.Volume)
		}
	}
	return countUpdate, priceUpdate, stockUpdate
}

func arrayToString(arr []int) string {
	str := ""
	for _, val := range arr {
		str += strconv.Itoa(val) + " "
	}
	return str
}

func (config CrawlingConfig) getSourcename() string {
	if config.Source.RightStufAnime {
		return "RightStufAnime"
	} else if config.Source.Amazon {
		return "Amazon"
	} else if config.Source.InStockTrades {
		return "InStockTrades"
	} else if config.Source.BookDepository {
		return "BookDepository"
	}
	return "Not Found"
}
