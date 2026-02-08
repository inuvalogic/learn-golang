package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"kasir-api/models"
	"kasir-api/services"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleTodayReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.GetTodayReport(w, r)
}

func (h *ReportHandler) HandleGetReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	h.GetReport(w, r)
}

func (h *ReportHandler) GetTodayReport(w http.ResponseWriter, r *http.Request) {
	nowUTC := time.Now().UTC()

	startOfDayUTC := time.Date(
		nowUTC.Year(),
		nowUTC.Month(),
		nowUTC.Day(),
		0, 0, 0, 0,
		time.UTC,
	)

	timerange := models.TimeRange{
		StartDate: startOfDayUTC,
		EndDate:   nowUTC,
	}

	report, err := h.service.GetReport(timerange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	startStr := r.URL.Query().Get("start_date")
	endStr := r.URL.Query().Get("end_date")

	if startStr == "" || endStr == "" {
		http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	// Parse date as UTC date (not local)
	startDate, err := time.ParseInLocation("2006-01-02", startStr, time.UTC)
	if err != nil {
		http.Error(w, "invalid start_date format", http.StatusBadRequest)
		return
	}

	endDate, err := time.ParseInLocation("2006-01-02", endStr, time.UTC)
	if err != nil {
		http.Error(w, "invalid end_date format", http.StatusBadRequest)
		return
	}

	startDate = time.Date(
		startDate.Year(), startDate.Month(), startDate.Day(),
		0, 0, 0, 0,
		time.UTC,
	)

	endDate = time.Date(
		endDate.Year(), endDate.Month(), endDate.Day(),
		23, 59, 59, 0,
		time.UTC,
	)

	timerange := models.TimeRange{
		StartDate: startDate,
		EndDate:   endDate,
	}

	report, err := h.service.GetReport(timerange)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(report)
}

