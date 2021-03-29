package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

func (c *Controller) getDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	userDataFound, err := c.DashboardRepo.GetDashboardByUserId(ctx, userId)
	if userDataFound == nil {
		utils.GetCORSErrResponse(w, "There's no data found for this user", http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	var data models.DashboardResponse

	data.FirstName = userDataFound.FirstName
	heartRate := utils.GetUnMarshalledHeartData(w, r, userDataFound.HeartRate)
	bloodPressure := utils.GetUnMarshalledBloodPressure(w, r, userDataFound.BloodPressure)
	bodyTemperature := utils.GetUnMarshalledBodyTemperature(w, r, userDataFound.BodyTemperature)
	glucose := utils.GetUnMarshalledGlucose(w, r, userDataFound.Glucose)
	oxygenSaturation := utils.GetUnMarshalledOxygenSaturation(w, r, userDataFound.OxygenSaturation)
	data.HeartRate = heartRate
	data.BloodPressure = bloodPressure
	data.BodyTemperature = bodyTemperature
	data.Glucose = glucose
	data.OxygenSaturation = oxygenSaturation
	data.AlertSent = userDataFound.AlertSent
	data.AlertsResponded = userDataFound.AlertsResponded

	err = json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) PostDashboard(w http.ResponseWriter, r *http.Request, requestData models.DashboardRequest) {
	// ctx := r.Context()
	// userId := ctx.Value("user_id").(string)
fmt.Println("New Data ==>", requestData)
	data := r.Body
	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if c.DashboardValidator(requestData) == false {
		utils.GetCORSErrResponse(w, "Fields are not properly formated", http.StatusBadRequest)
		return
	} else {
		// heartRate := utils.GetMarshalledHeartData(w, r, requestData.HeartRate)
		// bodyTemperature := utils.GetMarshalledBodyTemperatureData(w, r, requestData.BodyTemperature)
		// bloodPressure := utils.GetMarshalledBloodPressureData(w, r, requestData.BloodPressure)
		// glucose := utils.GetMarshalledGlucoseData(w, r, requestData.Glucose)
		// oxygenSaturation := utils.GetMarshalledOxygenSaturationData(w, r, requestData.OxygenSaturation)
		// record := models.Dashboard{
		// 	UserId:           userId,
		// 	HeartRate:        string(heartRate),
		// 	BloodPressure:    string(bloodPressure),
		// 	BodyTemperature:  string(bodyTemperature),
		// 	Glucose:          string(glucose),
		// 	OxygenSaturation: string(oxygenSaturation),
		// }

		// c.DashboardRepo.CreateOrUpdateRecordByUserId(ctx, userId, &record)
	}
}

