package controllers

import (
	"fmt"
	"net/http"
)

func (c *Controller) postLogin(w http.ResponseWriter, r *http.Request) {
	// user := &models.User{}
	w.Write([]byte("Login Successful!"))
}

func (c *Controller) postRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
}

func (c *Controller) getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pillshare-v1.0"))
}
