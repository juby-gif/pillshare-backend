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
	fmt.Println(medicalRecordFound)

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
	fmt.Println(requestData)
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
	intervals := requestData.Intervals
	reason := requestData.Reason

	fmt.Println(name, userId)
	fmt.Println(dose)
	fmt.Println(measure)
	fmt.Println(isDeleted)
	fmt.Println(dosage)
	fmt.Println(beforeOrAfter)
	fmt.Println(duration)
	fmt.Println(startDate)
	fmt.Println(endDate)
	fmt.Println(intervals)
	fmt.Println(reason)
	// record := models.MedicalRecord{
	// 	UserId:        userId,
	// 	Name:          firnamestName,
	// 	Dose:          dose,
	// 	Measure:       measure,
	// 	IsDeleted:     isDeleted,
	// 	Dosage:        dosage,
	// 	BeforeOrAfter: beforeOrAfter,
	// 	Duration:      duration,
	// 	StartDate:     startDate,
	// 	EndDate:       endDate,
	// 	Intervals:     intervals,
	// 	Reason:        reason,
	// }
	// fmt.Println(record)
	// c.MedicalRepo.CreateOrUpdateMedicalRecordByUserId(ctx, userId, &record)
}
