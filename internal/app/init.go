package app

import (
	"context"
	"luminnovel/internal/service/bookdepository"
	"luminnovel/internal/usecase/report"
	"net/http"
	"time"

	"luminnovel/internal/handler"
	productHandler "luminnovel/internal/handler/product"
	reportHandler "luminnovel/internal/handler/report"
	"luminnovel/internal/repository/googlesheet"
	httpRepo "luminnovel/internal/repository/http"
	"luminnovel/internal/repository/mongodb"
	"luminnovel/internal/service/crawling"
	"luminnovel/internal/service/rightstufanime"
	"luminnovel/internal/usecase/product"
)

func InitHTTP(ctx context.Context) {
	clientDB := initMongoDB(ctx)

	// get collection as ref
	database := clientDB.Database("LuminNovel")
	mongoRepo := mongodb.New(database)

	sheetService := initGoogleSheetAPI(ctx)
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	httpRepo := httpRepo.New(httpClient)
	googleSheetRepo := googlesheet.New(sheetService)

	bookDepoRsc := bookdepository.NewResource(httpRepo, mongoRepo)
	rightStufRsc := rightstufanime.NewResource(httpRepo, mongoRepo)
	crawlingRsc := crawling.NewResource(googleSheetRepo)

	bookDepoSvc := bookdepository.NewService(bookDepoRsc)
	rightStufSvc := rightstufanime.NewService(rightStufRsc)
	crawlingSvc := crawling.NewService(crawlingRsc)

	usecaseProduct := product.New(crawlingSvc, rightStufSvc, bookDepoSvc)
	usecaseReport := report.New()
	handlerProduct := productHandler.New(usecaseProduct)
	handlerReport := reportHandler.New(usecaseReport)

	handler.ServeHTTP(handlerProduct, handlerReport)
}
