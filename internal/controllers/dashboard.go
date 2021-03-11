package controllers

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func (c *Controller) postDashboard(w http.ResponseWriter, r *http.Request) {
	ctx:= r.Context()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	data:= r.Body
	var requestData models.DashboardRequest
	userId := ctx.Value("user_id").(string)
	fmt.Println(userId)

	err := json.NewDecoder(data).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(requestData.Params)
}