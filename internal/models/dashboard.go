package models

import (
	"context"
	// "time"
)

type HeartRate struct {
	InstrumentID int     `json:"instrumentId"`
	Reading      int     `json:"reading"`
	Percentage   float64 `json:"percentage"`
	Time         string  `json:"time"`
}

type Glucose struct {
	InstrumentID int     `json:"instrumentId"`
	Reading      int     `json:"reading"`
	Percentage   float64 `json:"percentage"`
	Time         string  `json:"time"`
}

type BloodPressure struct {
	InstrumentID    int     `json:"instrumentId"`
	SystoleReading  int     `json:"systoleReading"`
	DiastoleReading int     `json:"diastoleReading"`
	Percentage      float64 `json:"percentage"`
	Time            string  `json:"time"`
}

type BodyTemperature struct {
	InstrumentID int     `json:"instrumentId"`
	Reading      int     `json:"reading"`
	Percentage   float64 `json:"percentage"`
	Time         string  `json:"time"`
}

type OxygenSaturation struct {
	InstrumentID int     `json:"instrumentId"`
	Reading      int     `json:"reading"`
	Percentage   float64 `json:"percentage"`
	Time         string  `json:"time"`
}

// type HealthCheck struct {
// 	HealthStatus string    `json:"healthStatus"`
// 	Time         time.Time `json:"time"`
// }

type DashboardRequest struct {
	FirstName        string            `json:"firstName"`
	UserId           string            `json:"userId"`
	HeartRate        *HeartRate        `json:"heartRate"`
	BloodPressure    *BloodPressure    `json:"bloodPressure"`
	BodyTemperature  *BodyTemperature  `json:"bodyTemperature"`
	Glucose          *Glucose          `json:"glucose"`
	OxygenSaturation *OxygenSaturation `json:"oxygenSaturation"`
	AlertSent        int               `json:"alertsSent"`
	AlertsResponded  int               `json:"alertsResponded"`
}

type DashboardResponse struct {
	FirstName        string            `json:"firstName"`
	HeartRate        *HeartRate        `json:"heartRate"`
	BloodPressure    *BloodPressure    `json:"bloodPressure"`
	BodyTemperature  *BodyTemperature  `json:"bodyTemperature"`
	Glucose          *Glucose          `json:"glucose"`
	OxygenSaturation *OxygenSaturation `json:"oxygenSaturation"`
	AlertSent        int               `json:"alertsSent"`
	AlertsResponded  int               `json:"alertsResponded"`
}

type Dashboard struct {
	FirstName        string `json:"firstName"`
	UserId           string `json:"userId"`
	HeartRate        string `json:"heartRate"`
	BloodPressure    string `json:"bloodPressure"`
	BodyTemperature  string `json:"bodyTemperature"`
	Glucose          string `json:"glucose"`
	OxygenSaturation string `json:"oxygenSaturation"`
	AlertSent        int    `json:"alertsSent"`
	AlertsResponded  int    `json:"alertsResponded"`
}

type DashboardRepo interface {
	CreateNewDataRecord(ctx context.Context, m *Dashboard) error
	CreateNewAlertsRecord(ctx context.Context, m *Dashboard) error
	GetDashboardByUserId(ctx context.Context, user_id string) (*Dashboard, error)
	UpdateRecordByUserId(ctx context.Context, m *Dashboard) error
	UpdateAlertsByUserId(ctx context.Context, m *Dashboard) error
	CreateOrUpdateRecordByUserId(ctx context.Context, userId string, m *Dashboard) error
	CreateOrUpdateAlertsByUserId(ctx context.Context, userId string, m *Dashboard) error
}
