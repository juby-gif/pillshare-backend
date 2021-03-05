package controllers

import (
	"net/http"
)

func (c *Controller) getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pillshare-v1.0"))
}
