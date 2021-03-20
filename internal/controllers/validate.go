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
		return false
	case data.HeartRate == nil:
		return false
	case data.BloodPressure == nil:
		return false
	case data.BodyTemperature == nil:
		return false
	case data.Glucose == nil:
		return false
	case data.OxygenSaturation == nil:
		return false
	default:
		return true
	}
}

func (c *Controller) MedicalDataValidator(data models.MedicalDataRequest) bool {

	switch {
	case data.Name == "":
		fmt.Println("1")
		return false
	case data.Dose == "":
		fmt.Println("2")
		return false
	case data.Measure == "":
		fmt.Println("3")
		return false
	case data.Dosage == "":
		fmt.Println("4")
		return false
	case data.BeforeOrAfter == "":
		fmt.Println("5")
		return false
	case data.Duration == 0:
		fmt.Println("6")
		return false
	case len(data.Intervals.Part) == 0:
		fmt.Println("7")
		return false
	case len(data.Intervals.Time) == 0:
		fmt.Println("8")
		return false
	default:
		return true
	}
}