package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

func (c *Controller) getVitalsRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `heartRateRecordFound`
	vitalsRecordFound, err := c.VitalsRepo.GetVitalsRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	if vitalsRecordFound == nil {
		utils.GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// If the record exists, we have to unmarshal it with the data `vitalsRecordFound`.
	// Get the unmarshalled record and save it to `records`
	var records models.VitalsRecordResponse

	records.HeartRate = utils.GetUnMarshalledHeartArr(w, r, vitalsRecordFound.HeartRate)
	records.BloodPressure = utils.GetUnMarshalledBloodPressureArr(w, r, vitalsRecordFound.BloodPressure)
	records.BodyTemperature = utils.GetUnMarshalledBodyTemperatureArr(w, r, vitalsRecordFound.BodyTemperature)
	records.Glucose = utils.GetUnMarshalledGlucoseArr(w, r, vitalsRecordFound.Glucose)
	records.OxygenSaturation = utils.GetUnMarshalledOxygenSaturationArr(w, r, vitalsRecordFound.OxygenSaturation)
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getHeartRateRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `heartRateRecordFound`
	heartRateRecordFound, err := c.VitalsRepo.GetHeartRateRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	if heartRateRecordFound == "" {
		utils.GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// If the record exists, we have to unmarshal it with the data `heartRateRecordFound`.
	// Get the unmarshalled record and save it to `records`
	records := utils.GetUnMarshalledHeartArr(w, r, heartRateRecordFound)
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func (c *Controller) getBloodPressureRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `bloodPressureRecordFound`
	bloodPressureRecordFound, err := c.VitalsRepo.GetBloodPressureRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	if bloodPressureRecordFound == "" {
		utils.GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// If the record exists, we have to unmarshal it with the data `bloodPressureRecordFound`.
	// Get the unmarshalled record and save it to `records`
	records := utils.GetUnMarshalledBloodPressureArr(w, r, bloodPressureRecordFound)
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getBodyTemperatureRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `bodyTemperatureRecordFound`
	bodyTemperatureRecordFound, err := c.VitalsRepo.GetBodyTemperatureRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	if bodyTemperatureRecordFound == "" {
		utils.GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// If the record exists, we have to unmarshal it with the data `bodyTemperatureRecordFound`.
	// Get the unmarshalled record and save it to `records`
	records := utils.GetUnMarshalledBodyTemperatureArr(w, r, bodyTemperatureRecordFound)
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getGlucoseRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `glucoseRecordFound`
	glucoseRecordFound, err := c.VitalsRepo.GetGlucoseRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	if glucoseRecordFound == "" {
		utils.GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// If the record exists, we have to unmarshal it with the data `glucoseRecordFound`.
	// Get the unmarshalled record and save it to `records`
	records := utils.GetUnMarshalledGlucoseArr(w, r, glucoseRecordFound)
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) getOxygenSaturationRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `oxygenSaturationRecordFound`
	oxygenSaturationRecordFound, err := c.VitalsRepo.GetOxygenSaturationRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	if oxygenSaturationRecordFound == "" {
		utils.GetCORSErrResponse(w, "No Data Found", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// If the record exists, we have to unmarshal it with the data `oxygenSaturationRecordFound`.
	// Get the unmarshalled record and save it to `records`
	records := utils.GetUnMarshalledOxygenSaturationArr(w, r, oxygenSaturationRecordFound)
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

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `vitalsRecordFound`
	vitalsRecordFound, err := c.VitalsRepo.GetVitalsRecordByUserId(ctx, userId)

	// If there's no record found in the database, we have tp create it
	// Create an empty slice and append the new record to it
	if vitalsRecordFound == nil {
		newHeartRateRecord := []*models.HeartRateData{}
		newBloodPressureDataRecord := []*models.BloodPressureData{}
		newBodyTemperatureRecord := []*models.BodyTemperatureData{}
		newGlucoseRecord := []*models.GlucoseData{}
		newOxygenSaturationRecord := []*models.OxygenSaturationData{}

		newHeartRateRecord = append(newHeartRateRecord, requestData.HeartRate)
		newBloodPressureDataRecord = append(newBloodPressureDataRecord, requestData.BloodPressure)
		newBodyTemperatureRecord = append(newBodyTemperatureRecord, requestData.BodyTemperature)
		newGlucoseRecord = append(newGlucoseRecord, requestData.Glucose)
		newOxygenSaturationRecord = append(newOxygenSaturationRecord, requestData.OxygenSaturation)
		marshalledHeartRateRecord := utils.GetMarshalledHeartDataArr(w, r, newHeartRateRecord)
		marshalledBloodPressureRecord := utils.GetMarshalledBloodPressureDataArr(w, r, newBloodPressureDataRecord)
		marshalledBodyTemperatureRecord := utils.GetMarshalledBodyTemperatureDataArr(w, r, newBodyTemperatureRecord)
		marshalledGlucoseRecord := utils.GetMarshalledGlucoseDataArr(w, r, newGlucoseRecord)
		marshalledOxygenSaturationRecord := utils.GetMarshalledOxygenSaturationDataArr(w, r, newOxygenSaturationRecord)

		// Marshall the new record into []bytes and will save it to `marshalledNewRecord`
		// marshalledHeartRateNewRecord := utils.GetMarshalledVitalsRecord(w, r, newRecord)
		record := models.VitalsRecord{
			UserId:           userId,
			HeartRate:        string(marshalledHeartRateRecord),
			BloodPressure:    string(marshalledBloodPressureRecord),
			BodyTemperature:  string(marshalledBodyTemperatureRecord),
			Glucose:          string(marshalledGlucoseRecord),
			OxygenSaturation: string(marshalledOxygenSaturationRecord),
		}
		fmt.Println(record)

		// Invoke `CreateOrUpdateVitalsRecordByUserId` method with the created record
		// and `userId` which will create or update the record in Postgresql
		c.VitalsRepo.CreateOrUpdateVitalsRecordByUserId(ctx, userId, &record)
		message := &models.MedicalDataResponse{
			Message: "You have successfully added your vitals.",
		}
		if err := json.NewEncoder(w).Encode(&message); err != nil {
			utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else {

		// If there's record found in the database, we have to unmarshall it
		// Save the unmarshalled record to `records` and append the new record to it
		unmarshalledHeartRateRecord := utils.GetUnMarshalledHeartArr(w, r, vitalsRecordFound.HeartRate)
		unmarshalledBloodPressureRecord := utils.GetUnMarshalledBloodPressureArr(w, r, vitalsRecordFound.BloodPressure)
		unmarshalledBodyTemperatureRecord := utils.GetUnMarshalledBodyTemperatureArr(w, r, vitalsRecordFound.BodyTemperature)
		unmarshalledGlucoseRecord := utils.GetUnMarshalledGlucoseArr(w, r, vitalsRecordFound.Glucose)
		unmarshalledOxygenSaturationRecord := utils.GetUnMarshalledOxygenSaturationArr(w, r, vitalsRecordFound.OxygenSaturation)
		unmarshalledHeartRateRecord = append(unmarshalledHeartRateRecord, requestData.HeartRate)
		unmarshalledBloodPressureRecord = append(unmarshalledBloodPressureRecord, requestData.BloodPressure)
		unmarshalledBodyTemperatureRecord = append(unmarshalledBodyTemperatureRecord, requestData.BodyTemperature)
		unmarshalledGlucoseRecord = append(unmarshalledGlucoseRecord, requestData.Glucose)
		unmarshalledOxygenSaturationRecord = append(unmarshalledOxygenSaturationRecord, requestData.OxygenSaturation)

		marshalledHeartRateRecord := utils.GetMarshalledHeartDataArr(w, r, unmarshalledHeartRateRecord)
		marshalledBloodPressureRecord := utils.GetMarshalledBloodPressureDataArr(w, r, unmarshalledBloodPressureRecord)
		marshalledBodyTemperatureRecord := utils.GetMarshalledBodyTemperatureDataArr(w, r, unmarshalledBodyTemperatureRecord)
		marshalledGlucoseRecord := utils.GetMarshalledGlucoseDataArr(w, r, unmarshalledGlucoseRecord)
		marshalledOxygenSaturationRecord := utils.GetMarshalledOxygenSaturationDataArr(w, r, unmarshalledOxygenSaturationRecord)

		//Marshall the updated record and save it as []bytes in `marshalledRecord`
		record := models.VitalsRecord{
			UserId:           userId,
			HeartRate:        string(marshalledHeartRateRecord),
			BloodPressure:    string(marshalledBloodPressureRecord),
			BodyTemperature:  string(marshalledBodyTemperatureRecord),
			Glucose:          string(marshalledGlucoseRecord),
			OxygenSaturation: string(marshalledOxygenSaturationRecord),
		}

		fmt.Println(record)

		// Invoke `CreateOrUpdateVitalsRecordByUserId` method with the updated record
		// and `userId` which will create or update the record in Postgresql
		c.VitalsRepo.CreateOrUpdateVitalsRecordByUserId(ctx, userId, &record)
		message := &models.MedicalDataResponse{
			Message: "You have successfully added your vitals.",
		}
		if err := json.NewEncoder(w).Encode(&message); err != nil {
			utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	}
}
