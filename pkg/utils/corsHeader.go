package utils

import (
	"net/http"
	"encoding/json"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func GetCORSErrResponse(w http.ResponseWriter, message string, code int) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	responseData := &models.LoginErrorResponse{
		Message:message,
	}
	if err := json.NewEncoder(w).Encode(&responseData); err != nil {
		http.Error(w, err.Error(), code)
	return
	}
}

func EmptyTSDErrorHandler(w http.ResponseWriter, tsdRecord []*models.TimeSeriesRecord, err error) bool {

	if tsdRecord == nil {
		GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return true
	}
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return true
	}
	return false
}

func EmptyBloodPressureErrorHandler(w http.ResponseWriter, bloodPressureRecord []*models.BloodPressureRecord, err error) bool {

	if bloodPressureRecord == nil {
		GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return true
	}
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return true
	}
	return false
}

