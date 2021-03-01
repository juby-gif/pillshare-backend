package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func (c *Controller) postLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Successful!"))
}

func (c *Controller) postRegister(w http.ResponseWriter, r *http.Request) {
	data := r.Body
	var requestData models.RegisterRequest

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// user := &models.User{
	// 	User_id:        uuid.NewString(),
	// 	First_name:     requestData.FirstName,
	// 	Middle_name:    requestData.MiddleName,
	// 	Last_name:      requestData.LastName,
	// 	Username:       requestData.Username,
	// 	Email:          requestData.Email,
	// 	Password:       requestData.Password,
	// 	Checked_status: requestData.CheckedStatus,
	// }
}

func (c *Controller) getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pillshare-v1.0"))
}
