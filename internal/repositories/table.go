package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/juby-gif/pillshare-server/internal/models"
)

type MedicalRepo struct {
	db *sql.DB
}

func NewMedicalRepo(db *sql.DB) *MedicalRepo {
	return &MedicalRepo{
		db: db,
	}
}

func (med *MedicalRepo) CreateNewMedicalRecord(ctx context.Context, m *models.MedicalRecord) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO medical_database (user_id,record) VALUES ($1, $2)"
	stmt, err := med.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.UserId,
		m.Record,
	)
	return err
}

func (med *MedicalRepo) GetMedicalRecordByUserId(ctx context.Context, userId string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.MedicalRecord)
	query := "SELECT record FROM medical_database WHERE user_id = $1"
	err := med.db.QueryRowContext(ctx, query, userId).Scan(
		&m.Record,
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
	return m.Record, nil
}

func (med *MedicalRepo) UpdateMedicalRecordByUserId(ctx context.Context, m *models.MedicalRecord) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE medical_database SET record = $1  WHERE user_id = $2"
	stmt, err := med.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.Record,
		m.UserId,
	)
	return err
}

func (med *MedicalRepo) CheckIfMedicalRecordExistsByUserId(ctx context.Context, userId string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool

	query := `SELECT 1 FROM medical_database WHERE user_id = $1;`

	err := med.db.QueryRowContext(ctx, query, userId).Scan(&exists)
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

func (med *MedicalRepo) CreateOrUpdateMedicalRecordByUserId(ctx context.Context, userId string, m *models.MedicalRecord) error {
	exists, err := med.CheckIfMedicalRecordExistsByUserId(context.Background(), userId)
	if err != nil {
		return err
	}

	if exists { // CASE 1 OF 2: Update
		updateErr := med.UpdateMedicalRecordByUserId(ctx, m)
		if updateErr != nil {
			return updateErr
		}
	} else { // CASE 2 OF 2: Create
		createErr := med.CreateNewMedicalRecord(ctx, m)
		if createErr != nil {
			return createErr
		}
	}
	return nil
}
