package controllers

import (
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
		return false
	case data.Dose == "":
		return false
	case data.Measure == "":
		return false
	case data.Dosage == "":
		return false
	case data.BeforeOrAfter == "":
		return false
	case data.Duration == 0:
		return false
	case len(data.Intervals.Part) == 0:
		return false
	case len(data.Intervals.Time) == 0:
		return false
	default:
		return true
	}
}

func (c *Controller) VitalsDataValidator(data models.VitalsRecordRequest) bool {

	switch {
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
