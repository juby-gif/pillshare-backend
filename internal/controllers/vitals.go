package controllers

import (
	"encoding/json"
	// "fmt"
	// "strconv"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
	// "github.com/juby-gif/pillshare-server/pkg/utils"
)

// func (c *Controller) getVitalsRecord(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	userId := ctx.Value("user_id").(string)

// 	// Looking into the database to see if there's any record for the current user using `userId`
// 	// Save the response from database to `heartRateRecordFound`
// 	vitalsRecordFound, err := c.VitalsRepo.GetAllTimeSeriesRecordByUserId(ctx, userId)

// 	// If there's no record found in the database send the error.
// 	if vitalsRecordFound == nil {
// 		utils.GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
// 		return
// 	}
// 	if err != nil {
// 		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
// 	}
// 	// If the record exists, we have to unmarshal it with the data `vitalsRecordFound`.
// 	// Get the unmarshalled record and save it to `records`
// 	var records models.VitalsRecordResponse

// 	records.HeartRate = utils.GetUnMarshalledHeartArr(w, r, vitalsRecordFound.HeartRate)
// 	records.BloodPressure = utils.GetUnMarshalledBloodPressureArr(w, r, vitalsRecordFound.BloodPressure)
// 	records.BodyTemperature = utils.GetUnMarshalledBodyTemperatureArr(w, r, vitalsRecordFound.BodyTemperature)
// 	records.Glucose = utils.GetUnMarshalledGlucoseArr(w, r, vitalsRecordFound.Glucose)
// 	records.OxygenSaturation = utils.GetUnMarshalledOxygenSaturationArr(w, r, vitalsRecordFound.OxygenSaturation)
// 	if err := json.NewEncoder(w).Encode(&records); err != nil {
// 		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

func (c *Controller) postVitalsRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)
	// user := ctx.Value("user").(models.User)

	data := r.Body
	var requestData models.VitalsRecordRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// // Validate the requestData
	// // If any of the fields are missing, it will return false which will send error to the client
	// If all the fields are validated it will return true
	// if c.VitalsDataValidator(requestData) == false {
	// 	utils.GetCORSErrResponse(w, "Fields are not properly formated", http.StatusBadRequest)
	// 	return
	// } else {

	record := models.TimeSeriesRecord{
		UserId:       userId,
		InstrumentID: requestData.HeartRate.InstrumentID,
		Time:         requestData.HeartRate.Time,
		Reading:      requestData.HeartRate.Reading,
	}
	c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &record)

	record = models.TimeSeriesRecord{
		UserId:       userId,
		InstrumentID: requestData.BodyTemperature.InstrumentID,
		Time:         requestData.BodyTemperature.Time,
		Reading:      requestData.BodyTemperature.Reading,
	}
	c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &record)

	record = models.TimeSeriesRecord{
		UserId:       userId,
		InstrumentID: requestData.Glucose.InstrumentID,
		Time:         requestData.Glucose.Time,
		Reading:      requestData.Glucose.Reading,
	}
	c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &record)

	record = models.TimeSeriesRecord{
		UserId:       userId,
		InstrumentID: requestData.OxygenSaturation.InstrumentID,
		Time:         requestData.OxygenSaturation.Time,
		Reading:      requestData.OxygenSaturation.Reading,
	}
	c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &record)

	brecord := models.BloodPressureRecord{
		UserId:          userId,
		InstrumentID:    requestData.BloodPressure.InstrumentID,
		Time:            requestData.BloodPressure.Time,
		SystoleReading:  requestData.BloodPressure.SystoleReading,
		DiastoleReading: requestData.BloodPressure.DiastoleReading,
	}
	c.VitalsRepo.CreateNewBloodPressureRecord(ctx, &brecord)
	// }
}
