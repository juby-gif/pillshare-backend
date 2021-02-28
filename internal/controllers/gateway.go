package controllers

import (
	"net/http"
)

func (c *Controller) postLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Successful!"))
}

func (c *Controller) postRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register Successful!"))
}

func (c *Controller) getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pillshare-v1.0"))
}
