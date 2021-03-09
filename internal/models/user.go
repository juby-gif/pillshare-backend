package models

import (
	"context"
)

type User struct {
	First_name               string `json:"firstName"`
	Middle_name              string `json:"middleName"`
	Last_name                string `json:"lastName"`
	Username                 string `json:"username"`
	Email                    string `json:"email"`
	Password                 string `json:"password"`
	Checked_status           bool   `json:"checkedStatus"`
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
	Body_mass_index_value    string `json:"bodyMassIndexValue"`
	Blood_group              string `json:"bloodGroup"`
	Underlying_health_issues string `json:"underlyingHealthIssues"`
	Other_health_issues      string `json:"otherHealthIssues"`
	Images                   string `json:"images"`
	User_id                  string `json:"userId"`
}

type RegisterRequest struct {
	FirstName     string `json:"firstName"`
	MiddleName    string `json:"middleName"`
	LastName      string `json:"lastName"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	CheckedStatus bool   `json:"checkedStatus"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message      string `json:"message"`
	Length       int    `json:"length"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateNewUser(ctx context.Context, user_id string, first_name string, middle_name string, last_name string, username string, email string, password string, checked_status bool, age string, gender string, dob string, address string, city string, province string, country string, zip string, phone string, weight string, height string, bmi string, body_mass_index_value string, blood_group string, underlying_health_issues string, other_health_issues string, images string) error
}
