package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/internal/repositories"
	"github.com/juby-gif/pillshare-server/pkg/utils"
)

var middleName string
var password []byte

func (c *Controller) postLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := r.Body

	var requestData models.LoginRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the requestData
	// If any of the fields Email or Password is missing it will return false
	// If all the fields are validated it will return true
	if c.LoginValidator(requestData) == false {
		http.Error(w, "Fields are not properly formated", http.StatusBadRequest)
		return
	} else {
		userFound, err := c.UserRepo.GetUserByEmail(ctx, requestData.Email)
		if userFound == nil {
			http.Error(w, "This user does not match our records", http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		if utils.CompareHashedPassword(w, r, []byte(userFound.Password), []byte(requestData.Password)) == false {
			http.Error(w, "The password you entered is incorrect", http.StatusBadRequest)
			return
		} else {

			// Generates session id for the logged-in user
			// Access the `secretKey` from the `.env` file
			sessionToken := uuid.New().String()
			secretKey, err := ioutil.ReadFile(".env")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			// Generate JWT Token pair (`accessToken`,`refreshToken`)
			accessToken, refreshToken, err := utils.GenerateJWTTokenPair([]byte(secretKey), sessionToken)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Get the length of the User struct field
			// The length will return a value of type "int"
			length := utils.GetLengthOfUserField(userFound)

			var responseData = models.LoginResponse{
				Message:      "Success! You're Logging in",
				Length:       length,
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}
			err = json.NewEncoder(w).Encode(&responseData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}
}

func (c *Controller) postRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := r.Body
	var requestData models.RegisterRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the requestData
	// If any of the fields FirstName,LastName,Username,Email,Password and CheckedStatus is missing it will return false
	// If all the fields are validated it will return true
	if c.RegisterValidator(requestData) == false {
		http.Error(w, "Fields are not properly formated", http.StatusBadRequest)
		return
	} else {
		userFound, err := c.UserRepo.GetUserByEmail(ctx, requestData.Email)
		if userFound != nil {
			http.Error(w, "This Email already exists", http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, "Internal Error", http.StatusBadRequest)
			return
		}

		// Check the MiddleName of the new user
		// Assign "null" to the middleName if the field is blank
		// Assign the value "requestData.MiddleName" to middleName if the field has real value
		if requestData.MiddleName == "" {
			middleName = "null"
		} else {
			middleName = requestData.MiddleName
		}

		if requestData.Password != "" {
			password = utils.GenerateHashedPassword(w, r, requestData.Password)
		}

		ur := repositories.NewUserRepo(c.db)
		ur.CreateNewUser(
			ctx,
			uuid.NewString(),
			requestData.FirstName,
			middleName,
			requestData.LastName,
			requestData.Username,
			requestData.Email,
			string(password),
			requestData.CheckedStatus,
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
			"null",
		)

		var responseData = models.RegisterResponse{
			Message: "Congratulations! You are successfully registered!",
		}
		err = json.NewEncoder(w).Encode(&responseData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
