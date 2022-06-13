package product

import (
	"context"
	"fmt"
	"luminnovel/internal/service/bookdepository"
	"strconv"
	"time"

	"luminnovel/internal/entity"
	"luminnovel/internal/service/crawling"
	"luminnovel/internal/service/rightstufanime"
)

func (usecase *Usecase) CrawlingProductSeries(ctx context.Context, title entity.ProductTitle, source entity.SiteSource) error {
	if source == "" {
		return usecase.CrawlingAllSourceProduct(ctx, title)
	}

	var (
		payload      []crawling.CrawlingPayload
		sourceConfig crawling.CrawlingSource
	)

	switch source {
	case entity.BookDepository:
		products, err := usecase.bookDepoSvc.FetchProductDataByTitle(ctx, string(title))
		if err != nil {
			return err
		}
		payload = convertCrawlingPayloadBookDepository(products)
		sourceConfig = crawling.CrawlingSource{BookDepository: true}
	case entity.RightStufAnime:
		products, err := usecase.rightStufSvc.FetchProductDataByTitle(ctx, string(title))
		if err != nil {
			return err
		}
		payload = convertCrawlingPayloadRightStufAnime(products)
		sourceConfig = crawling.CrawlingSource{RightStufAnime: true}
	default:
		return nil
	}

	err := usecase.crawlingSheetSvc.UpdateCrawlingSheet(payload, crawling.CrawlingConfig{
		Series: title,
		Source: sourceConfig,
	})
	if err != nil {
		return err
	}
	return nil
}

func convertCrawlingPayloadRightStufAnime(products []rightstufanime.Product) []crawling.CrawlingPayload {
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

func convertCrawlingPayloadBookDepository(products []bookdepository.Product) []crawling.CrawlingPayload {
	payload := make([]crawling.CrawlingPayload, len(products))
	for index, product := range products {
		volume, _ := strconv.Atoi(product.Volume)
		payload[index] = crawling.CrawlingPayload{
			Volume:  volume,
			Price:   formatRupiah(product.Price),
			InStock: product.InStock,
		}
	}
	return payload
}

func (usecase *Usecase) CrawlingAllSourceProduct(ctx context.Context, title entity.ProductTitle) error {
	for _, source := range entity.ListAllSource {
		err := usecase.CrawlingProductSeries(ctx, title, source)
		if err != nil {
			return err
		}
	}
	return nil
}

func (usecase *Usecase) CrawlingAllProductSeries(ctx context.Context) error {
	for _, title := range entity.ListAllTitles {
		fmt.Printf("Start Crawling %s \n", title)
		err := usecase.CrawlingAllSourceProduct(ctx, title)
		if err != nil {
			return err
		}
		fmt.Printf("Success Crawling %s \n", title)
		time.Sleep(5 * time.Second)
	}
	return nil
}
