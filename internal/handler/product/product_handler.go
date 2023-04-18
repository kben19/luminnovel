package product

import (
	"context"
	"log"
	"luminnovel/internal/entity"
	"net/http"
	"time"
)

func (handler *Handler) HandleGetCrawlingProduct(w http.ResponseWriter, r *http.Request) {
	paramTitle := r.FormValue("title")
	if !validateTitle(paramTitle) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid title"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	paramSource := entity.SiteSource(r.FormValue("source"))
	if !paramSource.Validate() {
		paramSource = ""
	}

	err := handler.usecaseProduct.CrawlingProductSeries(ctx, entity.ProductTitle(paramTitle), paramSource)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Something wrong happened"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Success"))
}

func validateTitle(title string) bool {
	for _, value := range entity.ListAllTitles {
		if title == string(value) {
			return true
		}
	}
	return false
}

func (handler *Handler) HandleGetAllCrawlingProduct(w http.ResponseWriter, r *http.Request) {
	go func() {
		ctx := context.Background()

		err := handler.usecaseProduct.CrawlingAllProductSeries(ctx)
		if err != nil {
			log.Println("Something wrong happened")
			log.Println(err)
		}
		log.Println("Success All")
	}()

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Success"))
}
