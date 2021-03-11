package controllers

import (
	"net/http"
)

func (c *Controller) getUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User"))
}