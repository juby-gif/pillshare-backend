package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/juby-gif/pillshare-server/internal/models"
	"github.com/juby-gif/pillshare-server/internal/repositories"
)

func (c *Controller) postLogin(w http.ResponseWriter, r *http.Request) {
	data := r.Body
	var requestData models.LoginRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestData.Password)
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
	userFound, err := c.UserRepo.GetUserByEmail(ctx, requestData.Email)
	if userFound != nil {
		http.Error(w, "This Email already exists", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Internal Error", http.StatusBadRequest)
		return
	}
	ur := repositories.NewUserRepo(c.db)
	ur.CreateNewUser(
		ctx,
		uuid.NewString(),
		requestData.FirstName,
		requestData.MiddleName,
		requestData.LastName,
		requestData.Username,
		requestData.Email,
		requestData.Password,
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
		Message: "You have been registered successfully!",
	}
	err = json.NewEncoder(w).Encode(&responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
