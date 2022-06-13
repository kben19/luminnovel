package app

import (
	"context"
	"luminnovel/internal/service/bookdepository"
	"net/http"
	"time"

	"luminnovel/internal/handler"
	productHandler "luminnovel/internal/handler/product"
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
	handlerProduct := productHandler.New(usecaseProduct)

	handler.ServeHTTP(handlerProduct)
}
