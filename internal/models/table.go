package models

import (
	"context"
	"time"
)

type Intervals struct {
	Part []string `json:"part"`
	Time []string `json:"time"`
}

type MedicalRecord struct {
	UserId        string    `json:"userId"`
	Name          string    `json:"name"`
	Dose          string    `json:"dose"`
	Measure       string    `json:"measure"`
	IsDeleted     bool      `json:"isDeleted"`
	Dosage        string    `json:"dosage"`
	BeforeOrAfter string    `json:"beforeOrAfter"`
	Duration      int       `json:"duration"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Intervals     string    `json:"intervals"`
	Reason        string    `json:"reason"`
}

type MedicalDataRequest struct {
	UserId        string     `json:"userId"`
	Name          string     `json:"name"`
	Dose          string     `json:"dose"`
	Measure       string     `json:"measure"`
	IsDeleted     bool       `json:"isDeleted"`
	Dosage        string     `json:"dosage"`
	BeforeOrAfter string     `json:"beforeOrAfter"`
	Duration      int        `json:"duration"`
	StartDate     time.Time  `json:"startDate"`
	EndDate       time.Time  `json:"endDate"`
	Intervals     *Intervals `json:"intervals"`
	Reason        string     `json:"reason"`
}

type MedicalDataResponse struct {
	message string `json:"message"`
}

type Taken struct {
	timestamp time.Time `json:"timestamp"`
	name      string    `json:"name"`
}

type Missed struct {
	timestamp time.Time `json:"timestamp"`
	name      string    `json:"name"`
}

type MedicalRepo interface {
	CreateNewMedicalRecord(ctx context.Context, m *MedicalRecord) error
	GetMedicalRecordByUserId(ctx context.Context, userId string) (*MedicalRecord, error)
	UpdateMedicalRecordByUserId(ctx context.Context, m *MedicalRecord) error
	CreateOrUpdateMedicalRecordByUserId(ctx context.Context, userId string, m *MedicalRecord) error
}
