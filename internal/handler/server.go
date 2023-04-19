package handler

import (
	"luminnovel/internal/handler/product"
	"luminnovel/internal/handler/report"
	"net/http"
)

func ServeHTTP(handlerProduct *product.Handler, handlerReport *report.Handler) {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/crawling/product", handlerProduct.HandleGetCrawlingProduct)
	http.HandleFunc("/crawling/product/all", handlerProduct.HandleGetAllCrawlingProduct)
	http.HandleFunc("/reporting/calculate", handlerReport.HandleCalculateReport)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Welcome to Lumin Novel!"))
}
