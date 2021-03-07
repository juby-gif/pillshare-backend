package controllers

import (
	"net/http"
)

func (c *Controller) getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
