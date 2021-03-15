package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func (c *Controller) getNavHeader(w http.ResponseWriter, r *http.Request, loggedInUser models.User) {

	//For debugging purpose only
	// fmt.Println("User =>", loggedInUser)
	var navHeaderResponse = &models.NavHeaderResponse{
		Fname: loggedInUser.FirstName,
		Lname: loggedInUser.LastName,
	}

	err := json.NewEncoder(w).Encode(&navHeaderResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
