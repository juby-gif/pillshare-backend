package models

import (
	"context"
	"time"
)

type HeartRateData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      float64   `json:"reading"`
	Time         time.Time `json:"time"`
}

type GlucoseData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      float64   `json:"reading"`
	Time         time.Time `json:"time"`
}

type BloodPressureData struct {
	InstrumentID    int       `json:"instrumentId"`
	SystoleReading  float64   `json:"systoleReading"`
	DiastoleReading float64   `json:"diastoleReading"`
	Time            time.Time `json:"time"`
}

type BodyTemperatureData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      float64   `json:"reading"`
	Time         time.Time `json:"time"`
}

type OxygenSaturationData struct {
	InstrumentID int       `json:"instrumentId"`
	Reading      float64   `json:"reading"`
	Time         time.Time `json:"time"`
}

type VitalsRecordRequest struct {
	HeartRate        *HeartRateData        `json:"heartRate"`
	BloodPressure    *BloodPressureData    `json:"bloodPressure"`
	BodyTemperature  *BodyTemperatureData  `json:"bodyTemperature"`
	Glucose          *GlucoseData          `json:"glucose"`
	OxygenSaturation *OxygenSaturationData `json:"oxygenSaturation"`
}

type TimeSeriesRecord struct {
	UserId       string    `json:"userId"`
	InstrumentID int       `json:"instrumentId"`
	Time         time.Time `json:"time"`
	Reading      float64   `json:"reading"`
}

type BloodPressureRecord struct {
	UserId          string    `json:"userId"`
	InstrumentID    int       `json:"instrumentId"`
	Time            time.Time `json:"time"`
	SystoleReading  float64   `json:"systoleReading"`
	DiastoleReading float64   `json:"diastoleReading"`
}

type VitalsRecordResponse struct {
	HeartRate        []*HeartRateData        `json:"heartRate"`
	BloodPressure    []*BloodPressureData    `json:"bloodPressure"`
	BodyTemperature  []*BodyTemperatureData  `json:"bodyTemperature"`
	Glucose          []*GlucoseData          `json:"glucose"`
	OxygenSaturation []*OxygenSaturationData `json:"oxygenSaturation"`
}

type VitalsRepo interface {
	CreateNewTimeSeriesRecord(ctx context.Context, m *TimeSeriesRecord) error
	CreateNewBloodPressureRecord(ctx context.Context, m *BloodPressureRecord) error
	GetAllTimeSeriesRecordByUserId(ctx context.Context) ([]*TimeSeriesRecord, error)
}
