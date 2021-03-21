package models

import (
	"context"
	"time"
)

type HeartRateData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      int       `json:"reading"`
	Time         time.Time `json:"time"`
}

type GlucoseData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      int       `json:"reading"`
	Time         time.Time `json:"time"`
}

type BloodPressureData struct {
	InstrumentID    int       `json:"instrumentId"`
	SystoleReading  int       `json:"systoleReading"`
	DiastoleReading int       `json:"diastoleReading"`
	Time            time.Time `json:"time"`
}

type BodyTemperatureData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      int       `json:"reading"`
	Time         time.Time `json:"time"`
}

type OxygenSaturationData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      int       `json:"reading"`
	Time         time.Time `json:"time"`
}

type VitalsRecordRequest struct {
	UserId           string                `json:"userId"`
	HeartRate        *HeartRateData        `json:"heartRate"`
	BloodPressure    *BloodPressureData    `json:"bloodPressure"`
	BodyTemperature  *BodyTemperatureData  `json:"bodyTemperature"`
	Glucose          *GlucoseData          `json:"glucose"`
	OxygenSaturation *OxygenSaturationData `json:"oxygenSaturation"`
}

type VitalsRecord struct {
	UserId           string `json:"userId"`
	HeartRate        string `json:"heartRate"`
	BloodPressure    string `json:"bloodPressure"`
	BodyTemperature  string `json:"bodyTemperature"`
	Glucose          string `json:"glucose"`
	OxygenSaturation string `json:"oxygenSaturation"`
}

type VitalsRecordResponse struct {
	HeartRate        []*HeartRateData        `json:"heartRate"`
	BloodPressure    []*BloodPressureData    `json:"bloodPressure"`
	BodyTemperature  []*BodyTemperatureData  `json:"bodyTemperature"`
	Glucose          []*GlucoseData          `json:"glucose"`
	OxygenSaturation []*OxygenSaturationData `json:"oxygenSaturation"`
}

type VitalsRepo interface {
	CreateNewVitalsRecord(ctx context.Context, m *VitalsRecord) error
	GetVitalsRecordByUserId(ctx context.Context, userId string) (*VitalsRecord, error)
	GetHeartRateRecordByUserId(ctx context.Context, userId string) (string, error)
	GetBloodPressureRecordByUserId(ctx context.Context, userId string) (string, error)
	GetBodyTemperatureRecordByUserId(ctx context.Context, userId string) (string, error)
	GetGlucoseRecordByUserId(ctx context.Context, userId string) (string, error)
	GetOxygenSaturationRecordByUserId(ctx context.Context, userId string) (string, error)
	UpdateVitalsRecordByUserId(ctx context.Context, m *VitalsRecord) error
	CreateOrUpdateVitalsRecordByUserId(ctx context.Context, userId string, m *VitalsRecord) error
}
