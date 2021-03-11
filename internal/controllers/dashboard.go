package controllers

import (
	// "encoding/json"
	"net/http"
	"fmt"

	// "github.com/juby-gif/pillshare-server/internal/models"
)

func (c *Controller) getDashboard(w http.ResponseWriter, r *http.Request) {
	ctx:= r.Context()
	userId := ctx.Value("user_id").(string)
	fmt.Println(userId)
}