package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func (c *Controller) getUserProfile(w http.ResponseWriter, r *http.Request) {
	

	ctx := r.Context()
	loggedInUser := ctx.Value("user").(models.User)
	fmt.Println("User =>", loggedInUser)

	var responseData = models.User{
		FirstName:              loggedInUser.FirstName,
		MiddleName:             loggedInUser.MiddleName,
		LastName:               loggedInUser.LastName,
		Username:               loggedInUser.Username,
		Email:                  loggedInUser.Email,
		CheckedStatus:          loggedInUser.CheckedStatus,
		Age:                    loggedInUser.Age,
		Gender:                 loggedInUser.Gender,
		Dob:                    loggedInUser.Dob,
		Address:                loggedInUser.Address,
		City:                   loggedInUser.City,
		Province:               loggedInUser.Province,
		Country:                loggedInUser.Country,
		Zip:                    loggedInUser.Zip,
		Phone:                  loggedInUser.Phone,
		Weight:                 loggedInUser.Weight,
		Height:                 loggedInUser.Height,
		BMI:                    loggedInUser.BMI,
		BodyMassIndexValue:     loggedInUser.BodyMassIndexValue,
		BloodGroup:             loggedInUser.BloodGroup,
		UnderlyingHealthIssues: loggedInUser.UnderlyingHealthIssues,
		OtherHealthIssues:      loggedInUser.OtherHealthIssues,
		// Images:                 loggedInUser.Images,
	}
	err := json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func (c *Controller) patchUserProfile(w http.ResponseWriter, r *http.Request) {
	data := r.Body
	var requestData models.User

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestData)
}
