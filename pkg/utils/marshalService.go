package utils

import (
	"encoding/json"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func GetMarshalledHeartData(w http.ResponseWriter, r *http.Request, data *models.HeartRate) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledHeartDataArr(w http.ResponseWriter, r *http.Request, data []*models.HeartRateData) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledBloodPressureData(w http.ResponseWriter, r *http.Request, data *models.BloodPressure) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledBloodPressureDataArr(w http.ResponseWriter, r *http.Request, data []*models.BloodPressureData) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledBodyTemperatureData(w http.ResponseWriter, r *http.Request, data *models.BodyTemperature) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledBodyTemperatureDataArr(w http.ResponseWriter, r *http.Request, data []*models.BodyTemperatureData) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledGlucoseData(w http.ResponseWriter, r *http.Request, data *models.Glucose) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledGlucoseDataArr(w http.ResponseWriter, r *http.Request, data []*models.GlucoseData) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledOxygenSaturationData(w http.ResponseWriter, r *http.Request, data *models.OxygenSaturation) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledOxygenSaturationDataArr(w http.ResponseWriter, r *http.Request, data []*models.OxygenSaturationData) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}

func GetMarshalledMedicalRecord(w http.ResponseWriter, r *http.Request, data []*models.Record) []byte {
	marshalledData, err := json.Marshal(data)
	if err != nil {
		GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return marshalledData
}
