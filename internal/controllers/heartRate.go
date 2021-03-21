package controllers

import (
	"net/http"
)

func (c *Controller) getHeartRate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heart Rate"))
}