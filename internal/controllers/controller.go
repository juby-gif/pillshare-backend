package controllers

import (
	"net/http"
	"strings"
)

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) HandleRequests(w http.ResponseWriter, r *http.Request) {
	URL := strings.Split(r.URL.Path, "/")[1:]
	n := len(URL)

	switch {
	case n == 3 && URL[2] == "version" && r.Method == "GET":
		c.getVersion(w, r)
	case n == 3 && URL[2] == "login" && r.Method == "POST":
		c.postLogin(w, r)
	case n == 3 && URL[2] == "register" && r.Method == "POST":
		c.postRegister(w, r)
	}
}
