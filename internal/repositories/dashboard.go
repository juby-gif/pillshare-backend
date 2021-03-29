package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/juby-gif/pillshare-server/internal/models"
)

type DashboardRepo struct {
	db *sql.DB
}

func NewDashboardRepo(db *sql.DB) *DashboardRepo {
	return &DashboardRepo{
		db: db,
	}
}

func (r *DashboardRepo) CreateNewDataRecord(ctx context.Context, m *models.Dashboard) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO dashboard_dataset (user_id,heart_rate,blood_pressure,body_temperature,glucose,oxygen_saturation) VALUES ($1, $2, $3, $4, $5, $6)"

	stmt, err := r.db.PrepareContext(ctx, query)
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

func (r *DashboardRepo) CreateNewAlertsRecord(ctx context.Context, m *models.Dashboard) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO dashboard_dataset (alerts_sent,alerts_responded) VALUES ($1, $2)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.AlertSent,
		m.AlertsResponded,
	)
	return err
}

func (r *DashboardRepo) GetDashboardByUserId(ctx context.Context, userId string) (*models.Dashboard, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.Dashboard)
	query := "SELECT user_id,heart_rate,blood_pressure,body_temperature,glucose,oxygen_saturation,alerts_sent,alerts_responded FROM dashboard_dataset WHERE user_id = $1"
	err := r.db.QueryRowContext(ctx, query, userId).Scan(
		&m.UserId,
		&m.HeartRate,
		&m.BloodPressure,
		&m.BodyTemperature,
		&m.Glucose,
		&m.OxygenSaturation,
		&m.AlertSent,
		&m.AlertsResponded,
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

func (r *DashboardRepo) UpdateRecordByUserId(ctx context.Context, m *models.Dashboard) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE dashboard_dataset SET heart_rate = $1,blood_pressure = $2,body_temperature = $3,glucose = $4,oxygen_saturation = $5  WHERE user_id = $6"
	stmt, err := r.db.PrepareContext(ctx, query)
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

func (r *DashboardRepo) UpdateAlertsByUserId(ctx context.Context, m *models.Dashboard) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE dashboard_dataset SET alerts_sent = $1, alerts_responded= $2  WHERE user_id = $3"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.AlertSent,
		m.AlertsResponded,
		m.UserId,
	)
	return err
}

func (r *DashboardRepo) CheckIfUserRecordExistsByUserId(ctx context.Context, userId string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool

	query := `SELECT 1 FROM dashboard_dataset WHERE user_id = $1;`

	err := r.db.QueryRowContext(ctx, query, userId).Scan(&exists)
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

func (r *DashboardRepo) CreateOrUpdateRecordByUserId(ctx context.Context, userId string, m *models.Dashboard) error {
	exists, err := r.CheckIfUserRecordExistsByUserId(context.Background(), userId)
	if err != nil {
		return err
	}

	if exists { // CASE 1 OF 2: Update
		updateErr := r.UpdateRecordByUserId(ctx, m)
		if updateErr != nil {
			return updateErr
		}
	} else { // CASE 2 OF 2: Create
		createErr := r.CreateNewDataRecord(ctx, m)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}

func (r *DashboardRepo) CreateOrUpdateAlertsByUserId(ctx context.Context, userId string, m *models.Dashboard) error {
	exists, err := r.CheckIfUserRecordExistsByUserId(context.Background(), userId)
	if err != nil {
		return err
	}

	if exists { // CASE 1 OF 2: Update
		updateErr := r.UpdateAlertsByUserId(ctx, m)
		if updateErr != nil {
			return updateErr
		}
	} else { // CASE 2 OF 2: Create
		createErr := r.CreateNewAlertsRecord(ctx, m)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}
