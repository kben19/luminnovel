package report

import (
	"context"
	"luminnovel/internal/usecase/report"
)

type usecaseReportProvider interface {
	CalculateMonthlySummaryReporting(ctx context.Context, path string) (report.SummaryReport, error)
}

type Handler struct {
	usecaseReport usecaseReportProvider
}

func New(reportUC usecaseReportProvider) *Handler {
	return &Handler{reportUC}
}
