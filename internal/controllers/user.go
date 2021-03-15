package controllers

import (
	"context"
	"encoding/json"
	// "fmt"
	"net/http"
	"time"

	cache "github.com/go-redis/cache/v8"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

func (c *Controller) getUserProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	loggedInUser := ctx.Value("user").(models.User)

	//For debugging purpose only
	// fmt.Println("User =>", loggedInUser)

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
	ctx := r.Context()
	loggedInUser := ctx.Value("user").(models.User)
	data := r.Body
	sessionToken := ctx.Value("session_uuid").(string)

	// For Debugging purpose only
	// fmt.Println(sessionToken)

	var requestData models.User
	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println(requestData)

	length := utils.GetLengthOfUserField(&requestData)
	var middleName string
	if length >= 20 {
		if requestData.MiddleName == "" {
			middleName = ""
		} else {
			middleName = requestData.MiddleName
		}
		user := models.User{
			FirstName:              requestData.FirstName,
			MiddleName:             middleName,
			LastName:               requestData.LastName,
			Username:               requestData.Username,
			Email:                  requestData.Email,
			CheckedStatus:          requestData.CheckedStatus,
			Age:                    requestData.Age,
			Gender:                 requestData.Gender,
			Dob:                    requestData.Dob,
			Address:                requestData.Address,
			City:                   requestData.City,
			Province:               requestData.Province,
			Country:                requestData.Country,
			Zip:                    requestData.Zip,
			Phone:                  requestData.Phone,
			Weight:                 requestData.Weight,
			Height:                 requestData.Height,
			BMI:                    requestData.BMI,
			BodyMassIndexValue:     requestData.BodyMassIndexValue,
			BloodGroup:             requestData.BloodGroup,
			UnderlyingHealthIssues: requestData.UnderlyingHealthIssues,
			OtherHealthIssues:      requestData.OtherHealthIssues,
			UserId:                 loggedInUser.UserId,
		}
		err := c.UserRepo.UpdateUser(ctx, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		

		userFound, err := c.UserRepo.GetUserByEmail(ctx, requestData.Email)
		if userFound == nil {
			utils.GetCORSErrResponse(w, "This user does not match our records", http.StatusBadRequest)
			return
		}

		// Implemented redis-cache
		mycache := c.cache
		ctx := context.Background()
		key := sessionToken
		value := userFound

		//Deletes the user data stored in cache using `key`
		err = c.cache.Delete(ctx, key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the cache with `Key` as `sessionToken`
		// and the new record `Value` as `userFound`
		// Set the expiration for the cache as 3 days
		if err := mycache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   key,
			Value: value,
			TTL:   time.Hour * 72,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var responseData = models.PatchResponse{
			Message: "You have successfully saved your data!",
			Length:  length,
		}
		err = json.NewEncoder(w).Encode(&responseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		responseData := &models.PatchResponse{
			Message: "Sorry, you have to complete the required fields before moving forward",
			Length:  length,
		}
		if err := json.NewEncoder(w).Encode(&responseData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}
