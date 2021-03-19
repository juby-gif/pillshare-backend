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

	medicalRecordFound, err := c.MedicalRepo.GetMedicalRecordByUserId(ctx, userId)
	if medicalRecordFound == nil {
		utils.GetCORSErrResponse(w, "There's no data found for this user", http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Println(err.Error())
	}
	
	var data models.MedicalDataResponse

	data.Name = medicalRecordFound.Name
	data.Dose = medicalRecordFound.Dose
	data.Measure = medicalRecordFound.Measure
	data.IsDeleted = medicalRecordFound.IsDeleted
	data.Dosage = medicalRecordFound.Dosage
	data.BeforeOrAfter = medicalRecordFound.BeforeOrAfter
	data.Duration = medicalRecordFound.Duration
	data.StartDate = medicalRecordFound.StartDate
	data.EndDate = medicalRecordFound.EndDate
	data.Intervals = utils.GetUnMarshalledIntervals(w, r, medicalRecordFound.Intervals)
	data.Reason = medicalRecordFound.Reason
	err = json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	// if c.MedicalDataValidator(requestData) == false {
	// 	utils.GetCORSErrResponse(w, "Fields are not properly formated", http.StatusBadRequest)
	// 	return
	// } else {
	name := requestData.Name
	dose := requestData.Dose
	measure := requestData.Measure
	isDeleted := requestData.IsDeleted
	dosage := requestData.Dosage
	beforeOrAfter := requestData.BeforeOrAfter
	duration := requestData.Duration
	startDate := requestData.StartDate
	endDate := requestData.EndDate
	intervals := utils.GetMarshalledIntervals(w, r, requestData.Intervals)
	reason := requestData.Reason

	record := models.MedicalRecord{
		UserId:        userId,
		Name:          name,
		Dose:          dose,
		Measure:       measure,
		IsDeleted:     isDeleted,
		Dosage:        dosage,
		BeforeOrAfter: beforeOrAfter,
		Duration:      int(duration),
		StartDate:     startDate,
		EndDate:       endDate,
		Intervals:     string(intervals),
		Reason:        reason,
	}
	c.MedicalRepo.CreateOrUpdateMedicalRecordByUserId(ctx, userId, &record)
}
