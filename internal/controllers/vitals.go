package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

func (c *Controller) getVitalsRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId` and instrumentId `1`
	// Save the response from database to `heartRateRecordFound`
	heartRateRecordFound, err := c.VitalsRepo.GetTimeSeriesRecordByInstrumentIdandUserId(ctx, userId, 1)

	// If there's no record found in the database send the error.
	status := utils.EmptyTSDErrorHandler(w, heartRateRecordFound, err)
	if status == true {
		return
	}

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `bloodPressureRecordFound`
	bloodPressureRecordFound, err := c.VitalsRepo.GetBloodPressureRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	status = utils.EmptyBloodPressureErrorHandler(w, bloodPressureRecordFound, err)
	if status == true {
		return
	}

	// Looking into the database to see if there's any record for the current user using `userId` and instrumentId `3`
	// Save the response from database to `bodyTemperatureRecordFound`
	bodyTemperatureRecordFound, err := c.VitalsRepo.GetTimeSeriesRecordByInstrumentIdandUserId(ctx, userId, 3)

	// If there's no record found in the database send the error.
	status = utils.EmptyTSDErrorHandler(w, bodyTemperatureRecordFound, err)
	if status == true {
		return
	}

	// Looking into the database to see if there's any record for the current user using `userId` and instrumentId `4`
	// Save the response from database to `glucoseRecordFound`
	glucoseRecordFound, err := c.VitalsRepo.GetTimeSeriesRecordByInstrumentIdandUserId(ctx, userId, 4)

	// If there's no record found in the database send the error.
	status = utils.EmptyTSDErrorHandler(w, glucoseRecordFound, err)
	if status == true {
		return
	}

	// Looking into the database to see if there's any record for the current user using `userId` and instrumentId `5`
	// Save the response from database to `oxygenSaturationRecordFound`
	oxygenSaturationRecordFound, err := c.VitalsRepo.GetTimeSeriesRecordByInstrumentIdandUserId(ctx, userId, 5)

	// If there's no record found in the database send the error.
	status = utils.EmptyTSDErrorHandler(w, oxygenSaturationRecordFound, err)
	if status == true {
		return
	}

	records := models.VitalsRecordResponse{
		HeartRate:        heartRateRecordFound,
		BloodPressure:    bloodPressureRecordFound,
		BodyTemperature:  bodyTemperatureRecordFound,
		Glucose:          glucoseRecordFound,
		OxygenSaturation: oxygenSaturationRecordFound,
	}

	if err := json.NewEncoder(w).Encode(&records); err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) postVitalsRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	data := r.Body
	var requestData models.VitalsRecordRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the requestData
	// If any of the fields are missing, it will return false which will send error to the client
	// If all the fields are validated it will return true
	if c.VitalsDataValidator(requestData) == false {
		utils.GetCORSErrResponse(w, "Fields are not properly formated", http.StatusBadRequest)
		return
	} else {

		hrecord := models.TimeSeriesRecord{
			UserId:       userId,
			InstrumentID: requestData.HeartRate.InstrumentID,
			Time:         requestData.HeartRate.Time,
			Reading:      requestData.HeartRate.Reading,
		}
		c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &hrecord)

		btrecord := models.TimeSeriesRecord{
			UserId:       userId,
			InstrumentID: requestData.BodyTemperature.InstrumentID,
			Time:         requestData.BodyTemperature.Time,
			Reading:      requestData.BodyTemperature.Reading,
		}
		c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &btrecord)

		grecord := models.TimeSeriesRecord{
			UserId:       userId,
			InstrumentID: requestData.Glucose.InstrumentID,
			Time:         requestData.Glucose.Time,
			Reading:      requestData.Glucose.Reading,
		}
		c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &grecord)

		orecord := models.TimeSeriesRecord{
			UserId:       userId,
			InstrumentID: requestData.OxygenSaturation.InstrumentID,
			Time:         requestData.OxygenSaturation.Time,
			Reading:      requestData.OxygenSaturation.Reading,
		}
		c.VitalsRepo.CreateNewTimeSeriesRecord(ctx, &orecord)

		bprecord := models.BloodPressureRecord{
			UserId:          userId,
			InstrumentID:    requestData.BloodPressure.InstrumentID,
			Time:            requestData.BloodPressure.Time,
			SystoleReading:  requestData.BloodPressure.SystoleReading,
			DiastoleReading: requestData.BloodPressure.DiastoleReading,
		}
		c.VitalsRepo.CreateNewBloodPressureRecord(ctx, &bprecord)
	}
}
