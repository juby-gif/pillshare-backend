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

func (vr *VitalsRepo) CreateNewVitalsRecord(ctx context.Context, m *models.VitalsRecord) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO user_vitals_database (user_id,heart_rate,blood_pressure,body_temperature,glucose,oxygen_saturation) VALUES ($1, $2, $3, $4, $5, $6)"
	stmt, err := vr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.UserId,
		m.HeartRate,
		m.BloodPressure,
		m.BodyTemperature,
		m.Glucose,
		m.OxygenSaturation,
	)
	return err
}

func (vr *VitalsRepo) GetVitalsRecordByUserId(ctx context.Context, userId string) (*models.VitalsRecord, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.VitalsRecord)
	query := "SELECT heart_rate,blood_pressure,body_temperature,glucose,oxygen_saturation FROM user_vitals_database WHERE user_id = $1"
	err := vr.db.QueryRowContext(ctx, query, userId).Scan(
		&m.HeartRate,
		&m.BloodPressure,
		&m.BodyTemperature,
		&m.Glucose,
		&m.OxygenSaturation,
	)
	if err != nil {
		fmt.Println(err)
		// CASE 1 OF 2: Cannot find record with that userId.
		if err == sql.ErrNoRows {
			return nil, nil
		} else { // CASE 2 OF 2: All other errors.
			return nil, err
		}
	}
	return m, nil
}

func (vr *VitalsRepo) GetHeartRateRecordByUserId(ctx context.Context, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.VitalsRecord)
	query := "SELECT heart_rate FROM user_vitals_database WHERE user_id = $1"
	err := vr.db.QueryRowContext(ctx, query, userId).Scan(
		&m.HeartRate,
	)
	if err != nil {
		fmt.Println(err)
		// CASE 1 OF 2: Cannot find record with that userId.
		if err == sql.ErrNoRows {
			return "", nil
		} else { // CASE 2 OF 2: All other errors.
			return "", err
		}
	}
	return m.HeartRate, nil
}

func (vr *VitalsRepo) GetBloodPressureRecordByUserId(ctx context.Context, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.VitalsRecord)
	query := "SELECT blood_pressure FROM user_vitals_database WHERE user_id = $1"
	err := vr.db.QueryRowContext(ctx, query, userId).Scan(
		&m.BloodPressure,
	)
	if err != nil {
		fmt.Println(err)
		// CASE 1 OF 2: Cannot find record with that userId.
		if err == sql.ErrNoRows {
			return "", nil
		} else { // CASE 2 OF 2: All other errors.
			return "", err
		}
	}
	return m.BloodPressure, nil
}
func (vr *VitalsRepo) GetBodyTemperatureRecordByUserId(ctx context.Context, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.VitalsRecord)
	query := "SELECT body_temperature FROM user_vitals_database WHERE user_id = $1"
	err := vr.db.QueryRowContext(ctx, query, userId).Scan(
		&m.BodyTemperature,
	)
	if err != nil {
		fmt.Println(err)
		// CASE 1 OF 2: Cannot find record with that userId.
		if err == sql.ErrNoRows {
			return "", nil
		} else { // CASE 2 OF 2: All other errors.
			return "", err
		}
	}
	return m.BodyTemperature, nil
}

func (vr *VitalsRepo) GetGlucoseRecordByUserId(ctx context.Context, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.VitalsRecord)
	query := "SELECT glucose FROM user_vitals_database WHERE user_id = $1"
	err := vr.db.QueryRowContext(ctx, query, userId).Scan(
		&m.Glucose,
	)
	if err != nil {
		fmt.Println(err)
		// CASE 1 OF 2: Cannot find record with that userId.
		if err == sql.ErrNoRows {
			return "", nil
		} else { // CASE 2 OF 2: All other errors.
			return "", err
		}
	}
	return m.Glucose, nil
}

func (vr *VitalsRepo) GetOxygenSaturationRecordByUserId(ctx context.Context, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.VitalsRecord)
	query := "SELECT oxygen_saturation FROM user_vitals_database WHERE user_id = $1"
	err := vr.db.QueryRowContext(ctx, query, userId).Scan(
		&m.OxygenSaturation,
	)
	if err != nil {
		fmt.Println(err)
		// CASE 1 OF 2: Cannot find record with that userId.
		if err == sql.ErrNoRows {
			return "", nil
		} else { // CASE 2 OF 2: All other errors.
			return "", err
		}
	}
	return m.OxygenSaturation, nil
}

func (vr *VitalsRepo) UpdateVitalsRecordByUserId(ctx context.Context, m *models.VitalsRecord) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE user_vitals_database SET heart_rate = $1,blood_pressure = $2,body_temperature = $3,glucose = $4,oxygen_saturation = $5 WHERE user_id = $6"
	stmt, err := vr.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.HeartRate,
		m.BloodPressure,
		m.BodyTemperature,
		m.Glucose,
		m.OxygenSaturation,
		m.UserId,
	)
	return err
}

func (vr *VitalsRepo) CheckIfVitalsRecordExistsByUserId(ctx context.Context, userId string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool

	query := `SELECT 1 FROM user_vitals_database WHERE user_id = $1;`

	err := vr.db.QueryRowContext(ctx, query, userId).Scan(&exists)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return false, nil
		} else { // CASE 2 OF 2: All other errors.
			return false, err
		}
	}
	return exists, nil
}

func (vr *VitalsRepo) CreateOrUpdateVitalsRecordByUserId(ctx context.Context, userId string, m *models.VitalsRecord) error {
	exists, err := vr.CheckIfVitalsRecordExistsByUserId(context.Background(), userId)
	if err != nil {
		return err
	}

	if exists { // CASE 1 OF 2: Update
		updateErr := vr.UpdateVitalsRecordByUserId(ctx, m)
		if updateErr != nil {
			return updateErr
		}
	} else { // CASE 2 OF 2: Create
		createErr := vr.CreateNewVitalsRecord(ctx, m)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}
