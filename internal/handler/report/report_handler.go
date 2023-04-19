package report

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (handler *Handler) HandleCalculateReport(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pathParam := r.FormValue("path")
	if pathParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid path"))
		return
	}
	pathCommissionParam := r.FormValue("path_commission")

	summary, err := handler.usecaseReport.CalculateMonthlySummaryReporting(ctx, pathParam, pathCommissionParam)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Something wrong happened"))
		return
	}

	resp, err := json.Marshal(summary)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Something wrong happened"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
