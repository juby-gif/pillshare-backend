package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/juby-gif/pillshare-server/internal/models"
)

type VitalsRepo struct {
	db *sql.DB
}

func NewVitalsRepo(db *sql.DB) *VitalsRepo {
	return &VitalsRepo{
		db: db,
	}
}

func (vr *VitalsRepo) CreateNewTimeSeriesRecord(ctx context.Context, m *models.TimeSeriesRecord) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO time_series_record_database (user_id,instrument_id,time,reading) VALUES ($1, $2, $3, $4)"
	stmt, err := vr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.UserId,
		m.InstrumentID,
		m.Time,
		m.Reading,
	)
	return err
}

func (vr *VitalsRepo) CreateNewBloodPressureRecord(ctx context.Context, m *models.BloodPressureRecord) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO blood_pressure_database (user_id,instrument_id,time,systole_reading,diastole_reading) VALUES ($1, $2, $3, $4, $5)"
	stmt, err := vr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.UserId,
		m.InstrumentID,
		m.Time,
		m.SystoleReading,
		m.DiastoleReading,
	)
	return err
}

func (r *VitalsRepo) GetTimeSeriesRecordByInstrumentIdandUserId(ctx context.Context, userId string, instrumentId int) ([]*models.TimeSeriesRecord, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT instrument_id,time,reading FROM time_series_record_database WHERE user_id = $1 AND instrument_id = $2"

	rows, err := r.db.QueryContext(ctx, query, userId, instrumentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var s []*models.TimeSeriesRecord
	for rows.Next() {
		m := new(models.TimeSeriesRecord)
		err = rows.Scan(
			&m.InstrumentID,
			&m.Time,
			&m.Reading,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		s = append(s, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return s, err
}

func (r *VitalsRepo) GetBloodPressureRecordByUserId(ctx context.Context, userId string) ([]*models.BloodPressureRecord, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "SELECT instrument_id,time,systole_reading,diastole_reading FROM blood_pressure_database WHERE user_id = $1"

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var s []*models.BloodPressureRecord
	for rows.Next() {
		m := new(models.BloodPressureRecord)
		err = rows.Scan(
			&m.InstrumentID,
			&m.Time,
			&m.SystoleReading,
			&m.DiastoleReading,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		s = append(s, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return s, err
}
