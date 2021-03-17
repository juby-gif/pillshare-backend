package controllers

import (
	"fmt"

	"github.com/juby-gif/pillshare-server/internal/models"
)

func (c *Controller) LoginValidator(data models.LoginRequest) bool {
	switch {
	case data.Email == "":
		return false
	case data.Password == "":
		return false
	default:
		return true
	}
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

func (c *Controller) DashboardValidator(data models.DashboardRequest) bool {

	switch {
	case data.FirstName == "":
		fmt.Println("1")
		return false
	case data.HeartRate == nil:
		fmt.Println("2")
		return false
	case data.BloodPressure == nil:
		fmt.Println("3")
		return false
	case data.BodyTemperature == nil:
		fmt.Println("4")
		return false
	case data.Glucose == nil:
		fmt.Println("5")
		return false
	case data.OxygenSaturation == nil:
		fmt.Println("6")
		return false
	default:
		return true
	}
}
