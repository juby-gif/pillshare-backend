package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

func (c *Controller) getMedicalDatum(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// Looking into the database to see if there's any record for the current user using `userId`
	// Save the response from database to `medicalRecordFound`
	medicalRecordFound, err := c.MedicalRepo.GetMedicalRecordByUserId(ctx, userId)

	// If there's no record found in the database send the error.
	if medicalRecordFound == "" {
		utils.GetCORSErrResponse(w, "There's no data found for this user", http.StatusBadRequest)
		return
	}
	if err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// If the record exists, we have to unmarshal it with the data `medicalRecordFound`.
	// Get the unmarshalled record and save it to `records`
	records := utils.GetUnMarshalledMedicalRecord(w, r, medicalRecordFound)
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *Controller) postMedicalRecord(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	data := r.Body
	var requestData models.MedicalDataRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the requestData
	// If any of the fields are missing, it will return false which will send error to the client
	// If all the fields are validated it will return true
	if c.MedicalDataValidator(requestData) == false {
		utils.GetCORSErrResponse(w, "Fields are not properly formated", http.StatusBadRequest)
		return
	} else {

		// Looking into the database to see if there's any record for the current user using `userId`
		// Save the response from database to `medicalRecordFound`
		medicalRecordFound, err := c.MedicalRepo.GetMedicalRecordByUserId(ctx, userId)

		// If there's no record found in the database, we have tp create it
		// Create an empty slice and append the new record to it
		if medicalRecordFound == "" {
			newRecord := []*models.Record{}
			newRecord = append(newRecord, &models.Record{
				Name:          requestData.Name,
				Dose:          requestData.Dose,
				Measure:       requestData.Measure,
				IsDeleted:     requestData.IsDeleted,
				Dosage:        requestData.Dosage,
				BeforeOrAfter: requestData.BeforeOrAfter,
				Duration:      requestData.Duration,
				StartDate:     requestData.StartDate,
				EndDate:       requestData.EndDate,
				Intervals:     requestData.Intervals,
				Reason:        requestData.Reason,
			})

			// Marshall the new record into []bytes and will save it to `marshalledNewRecord`
			marshalledNewRecord := utils.GetMarshalledMedicalRecord(w, r, newRecord)
			record := models.MedicalRecord{
				UserId: userId,
				Record: string(marshalledNewRecord),
			}

			// Invoke `CreateOrUpdateMedicalRecordByUserId` method with the created record
			// and `userId` which will create or update the record in Postgresql
			c.MedicalRepo.CreateOrUpdateMedicalRecordByUserId(ctx, userId, &record)
			message := &models.MedicalDataResponse{
				Message: "You have successfully added your pill.",
			}
			if err := json.NewEncoder(w).Encode(&message); err != nil {
				utils.GetCORSErrResponse(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {

			// If there's record found in the database, we have to unmarshall it
			// Save the unmarshalled record to `records` and append the new record to it
			records := utils.GetUnMarshalledMedicalRecord(w, r, medicalRecordFound)
			records = append(records, &models.Record{
				Name:          requestData.Name,
				Dose:          requestData.Dose,
				Measure:       requestData.Measure,
				IsDeleted:     requestData.IsDeleted,
				Dosage:        requestData.Dosage,
				BeforeOrAfter: requestData.BeforeOrAfter,
				Duration:      requestData.Duration,
				StartDate:     requestData.StartDate,
				EndDate:       requestData.EndDate,
				Intervals:     requestData.Intervals,
				Reason:        requestData.Reason,
			})

			//Marshall the updated record and save it as []bytes in `marshalledRecord`
			marshalledRecord := utils.GetMarshalledMedicalRecord(w, r, records)
			record := models.MedicalRecord{
				UserId: userId,
				Record: string(marshalledRecord),
			}

			// Invoke `CreateOrUpdateMedicalRecordByUserId` method with the updated record
			// and `userId` which will create or update the record in Postgresql
			c.MedicalRepo.CreateOrUpdateMedicalRecordByUserId(ctx, userId, &record)
			message := &models.MedicalDataResponse{
				Message: "You have successfully added your pill.",
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
