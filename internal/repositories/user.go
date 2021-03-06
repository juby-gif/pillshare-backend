package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/juby-gif/pillshare-server/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateNewUser(ctx context.Context, user_id string, first_name string, middle_name string, last_name string, username string, email string, password string, checked_status bool, age string, gender string, dob string, address string, city string, province string, country string, zip string, phone string, weight string, height string, bmi string, body_mass_index_value string, blood_group string, underlying_health_issues string, other_health_issues string, images string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	query := "INSERT INTO users (user_id, first_name,middle_name,last_name,username,email,password,checked_status,age,gender,dob,address,city,province,country,zip,phone,weight,height,bmi,body_mass_index_value,blood_group,underlying_health_issues,other_health_issues,images) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		user_id,
		first_name,
		middle_name,
		last_name,
		username,
		email,
		password,
		checked_status,
		age,
		gender,
		dob,
		address,
		city,
		province,
		country,
		zip,
		phone,
		weight,
		height,
		bmi,
		body_mass_index_value,
		blood_group,
		underlying_health_issues,
		other_health_issues,
		images,
	)
	return err
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.User)

	query := "SELECT user_id, first_name,middle_name,last_name,username,email,password,checked_status,age,gender,dob,address,city,province,country,zip,phone,weight,height,bmi,body_mass_index_value,blood_group,underlying_health_issues,other_health_issues,images FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&m.User_id,
		&m.First_name,
		&m.Middle_name,
		&m.Last_name,
		&m.Username,
		&m.Email,
		&m.Password,
		&m.Checked_status,
		&m.Age,
		&m.Gender,
		&m.Dob,
		&m.Address,
		&m.City,
		&m.Province,
		&m.Country,
		&m.Zip,
		&m.Phone,
		&m.Weight,
		&m.Height,
		&m.BMI,
		&m.Body_mass_index_value,
		&m.Blood_group,
		&m.Underlying_health_issues,
		&m.Other_health_issues,
		&m.Images,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return nil, nil
		} else { // CASE 2 OF 2: All other errors.
			return nil, err
		}
	}
	return m, nil
}
