package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

func (c *Controller) getMedicalDatum(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	// medicalRecordFound, err := c.MedicalRepo.GetMedicalRecordByUserId(ctx, userId)
	// if medicalRecordFound == nil {
	// 	utils.GetCORSErrResponse(w, "There's no data found for this user", http.StatusBadRequest)
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	fmt.Println(userId)

	// var data models.MedicalDataResponse

	// data.Name = medicalRecordFound.Name
	// data.Dose = medicalRecordFound.Dose
	// data.Measure = medicalRecordFound.Measure
	// data.IsDeleted = medicalRecordFound.IsDeleted
	// data.Dosage = medicalRecordFound.Dosage
	// data.BeforeOrAfter = medicalRecordFound.BeforeOrAfter
	// data.Duration = medicalRecordFound.Duration
	// data.StartDate = medicalRecordFound.StartDate
	// data.EndDate = medicalRecordFound.EndDate
	// data.Intervals = utils.GetUnMarshalledIntervals(w, r, medicalRecordFound.Intervals)
	// data.Reason = medicalRecordFound.Reason
	// err = json.NewEncoder(w).Encode(&data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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
	if c.MedicalDataValidator(requestData) == false {
		utils.GetCORSErrResponse(w, "Fields are not properly formated", http.StatusBadRequest)
		return
	} else {
		medicalRecordFound, err := c.MedicalRepo.GetMedicalRecordByUserId(ctx, userId)
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
			marshalledNewRecord := utils.GetMarshalledMedicalRecord(w, r, newRecord)
			record := models.MedicalRecord{
				UserId: userId,
				Record: string(marshalledNewRecord),
			}
			fmt.Println(record)

			c.MedicalRepo.CreateOrUpdateMedicalRecordByUserId(ctx, userId, &record)
		} else {
			records := utils.GetUnMarshalledMedicalRecord(w, r, medicalRecordFound)
			fmt.Println("Old record---> ", records)
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
			marshalledRecord := utils.GetMarshalledMedicalRecord(w, r, records)
			record := models.MedicalRecord{
				UserId: userId,
				Record: string(marshalledRecord),
			}
			fmt.Println(record)

			c.MedicalRepo.CreateOrUpdateMedicalRecordByUserId(ctx, userId, &record)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
