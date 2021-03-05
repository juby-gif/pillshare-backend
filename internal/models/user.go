package models

import (
	"context"
)

type User struct {
	First_name               string `json:"first_name"`
	Middle_name              string `json:"middle_name"`
	Last_name                string `json:"last_name"`
	Username                 string `json:"username"`
	Email                    string `json:"email"`
	Password                 string `json:"password"`
	Checked_status           bool   `json:"checked_status"`
	Age                      string `json:"age"`
	Gender                   string `json:"gender"`
	Dob                      string `json:"dob"`
	Address                  string `json:"address"`
	City                     string `json:"city"`
	Province                 string `json:"province"`
	Country                  string `json:"country"`
	Zip                      string `json:"zip"`
	Phone                    string `json:"phone"`
	Weight                   string `json:"weight"`
	Height                   string `json:"height"`
	BMI                      string `json:"bmi"`
	Body_mass_index_value    string `json:"body_mass_index_value"`
	Blood_group              string `json:"blood_group"`
	Underlying_health_issues string `json:"underlying_health_issues"`
	Other_health_issues      string `json:"other_health_issues"`
	Images                   string `json:"images"`
	User_id                  string `json:"user_id"`
}

type RegisterRequest struct {
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	LastName      string `json:"last_name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	CheckedStatus bool   `json:"checked_status"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateNewUser(ctx context.Context, user_id string, first_name string, middle_name string, last_name string, username string, email string, password string, checked_status bool, age string, gender string, dob string, address string, city string, province string, country string, zip string, phone string, weight string, height string, bmi string, body_mass_index_value string, blood_group string, underlying_health_issues string, other_health_issues string, images string) error
}
