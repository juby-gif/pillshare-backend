package models

import (
	"time"
)

type HeartRate struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      int       `json:"reading"`
	Percentage   float64   `json:"percentage"`
	Time         time.Time `json:"time"`
}

type Glucose struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      int       `json:"reading"`
	Percentage   float64   `json:"percentage"`
	Time         time.Time `json:"time"`
}

type BloodPressure struct {
	InstrumentID    int       `json:"instrumentId"`
	SystoleReading  int       `json:"systoleReading"`
	DiastoleReading int       `json:"diastoleReading"`
	Percentage      float64   `json:"percentage"`
	Time            time.Time `json:"time"`
}

type OxygenSaturation struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      int       `json:"reading"`
	Percentage   float64   `json:"percentage"`
	Time         time.Time `json:"time"`
}

type HealthCheck struct {
	HealthStatus string    `json:"healthStatus"`
	Time         time.Time `json:"time"`
}

type DashboardResponse struct {
	FirstName           string            `json:"firstName"`
	UserId              string            `json:"userId"`
	HeartRateOBJ        *HeartRate        `json:"heartRateOBJ"`
	BloodPressureOBJ    *BloodPressure    `json:"bloodPressureOBJ"`
	GlucoseOBJ          *Glucose          `json:"glucoseOBJ"`
	OxygenSaturationOBJ *OxygenSaturation `json:"oxygenSaturationOBJ"`
	HealthCheckOBJ      *HealthCheck      `json:"healthCheckOBJ"`
	AlertSent           int               `json:"alertSent"`
	AlertsResponded     int               `json:"alertResponded"`
}

