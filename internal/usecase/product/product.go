package product

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"luminnovel/internal/entity"
	"luminnovel/internal/service/crawling"
	"luminnovel/internal/service/rightstufanime"
)

func (usecase *Usecase) CrawlingProductSeries(ctx context.Context, title entity.ProductTitle) error {
	// right stuf anime
	products, err := usecase.rightStufSvc.FetchProductDataByTitle(ctx, string(title))
	if err != nil {
		return err
	}

	payload := convertCrawlingPayload(products)
	err = usecase.crawlingSheetSvc.UpdateCrawlingSheet(payload, crawling.CrawlingConfig{
		Series: title,
		Source: crawling.CrawlingSource{RightStufAnime: true},
	})
	if err != nil {
		return err
	}
	return nil
}

func convertCrawlingPayload(products []rightstufanime.Product) []crawling.CrawlingPayload {
	payload := make([]crawling.CrawlingPayload, len(products))
	for index, product := range products {
		volume, _ := strconv.Atoi(product.Volume)
		payload[index] = crawling.CrawlingPayload{
			Volume:   volume,
			Price:    product.PriceFmt,
			InStock:  product.InStock,
			PreOrder: product.PreOrder,
		}
	}
	return payload
}

func (usecase *Usecase) CrawlingAllProductSeries(ctx context.Context) error {
	for _, title := range entity.ListAllTitles {
		fmt.Printf("Start Crawling %s \n", title)
		err := usecase.CrawlingProductSeries(ctx, title)
		if err != nil {
			return err
		}
		fmt.Printf("Success Crawling %s \n", title)
		time.Sleep(5 * time.Second)
	}
	return nil
}
