package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	cache "github.com/go-redis/cache/v8"
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

			// Implemented redis-cache
			mycache := c.cache

			// Set the cache with `Key` as `sessionToken`
			// and `Value` as `userFound`
			// Set the expiration for the cache as 3 days
			ctx := context.Background()
			key := sessionToken
			value := userFound
			if err := mycache.Set(&cache.Item{
				Ctx:   ctx,
				Key:   key,
				Value: value,
				TTL:   time.Hour * 72,
			}); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Get the length of the User struct field
			// The length will return a value of type "int"
			length := utils.GetLengthOfUserField(userFound)

			var responseData = &models.LoginResponse{
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

func (c *Controller) postRefreshToken(w http.ResponseWriter, r *http.Request, accessToken string) {

	// Access the `secretKey` from the `.env` file
	secretKey, err := ioutil.ReadFile(".env")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Verify our refresh token.
	// Pass `secretKey` and `accessToken` as the parameters for `ProcessJWTToken`
	// Returns `sessionToken` if the `err` is nil or
	// Return `Unauthorized - refresh token expired or invalid` is there is error
	sessionToken, err := utils.ProcessJWTToken([]byte(secretKey), accessToken)
	if err != nil {
		http.Error(w, "Unauthorized - refresh token expired or invalid", http.StatusUnauthorized)
		return
	}

	// Generate our JWT token.
	accessToken, refreshToken, err := utils.GenerateJWTTokenPair([]byte(secretKey), sessionToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generates `responseData` with `AccessToken` and `RefreshToken`
	responseData := models.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := json.NewEncoder(w).Encode(&responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// curl -X POST -H "Authorization:JWT eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUwOTIwOTMsInNlc3Npb25fdXVpZCI6IjIyYmQ4ZjhjLTI5MTAtNGY4NC05NDQ3LTI5ZWY3OTczODUxNyJ9.IGMsdg1HYII1xQxOuw0S6GaBpFM63QUHq62iv73BnOw" -H "Content-type:application/json" -H "Accept: application/json" -d '{"username":"lalla","password":"123pass"}' http://127.0.0.1:5000/api/v1/login
