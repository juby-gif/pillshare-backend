package controllers

import (
	"io"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func (c *Controller) LoginValidator(data io.ReadCloser) bool {
	return true
}

func (c *Controller) RegisterValidator(data models.RegisterRequest) bool {

	switch {
	case data.FirstName == "":
		return false
	case data.LastName == "":
		return false
	case data.Username == "":
		return false
	case data.Email == "":
		return false
	case data.Password == "":
		return false
	default:
		return true
	}
}
