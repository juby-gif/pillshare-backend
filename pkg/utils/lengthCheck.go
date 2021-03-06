package utils

import (
	"github.com/juby-gif/pillshare-server/internal/models"
)

func GetLengthOfUserField(user *models.User) int {
	userFieldsArr := []interface{}{}
	if user.First_name != "null" {
		userFieldsArr = append(userFieldsArr, user.First_name)
	}
	if user.Middle_name != "null" {
		userFieldsArr = append(userFieldsArr, user.Middle_name)
	}
	if user.Last_name != "null" {
		userFieldsArr = append(userFieldsArr, user.Last_name)
	}
	if user.Username != "null" {
		userFieldsArr = append(userFieldsArr, user.Username)
	}
	if user.Email != "null" {
		userFieldsArr = append(userFieldsArr, user.Email)
	}
	if user.Password != "null" {
		userFieldsArr = append(userFieldsArr, user.Password)
	}

	userFieldsArr = append(userFieldsArr, user.Checked_status)

	if user.Age != "null" {
		userFieldsArr = append(userFieldsArr, user.Age)
	}
	if user.Gender != "null" {
		userFieldsArr = append(userFieldsArr, user.Gender)
	}
	if user.Dob != "null" {
		userFieldsArr = append(userFieldsArr, user.Dob)
	}
	if user.Address != "null" {
		userFieldsArr = append(userFieldsArr, user.Address)
	}
	if user.City != "null" {
		userFieldsArr = append(userFieldsArr, user.City)
	}
	if user.Country != "null" {
		userFieldsArr = append(userFieldsArr, user.Country)
	}
	if user.Province != "null" {
		userFieldsArr = append(userFieldsArr, user.Province)
	}
	if user.Zip != "null" {
		userFieldsArr = append(userFieldsArr, user.Zip)
	}
	if user.Phone != "null" {
		userFieldsArr = append(userFieldsArr, user.Phone)
	}
	if user.Weight != "null" {
		userFieldsArr = append(userFieldsArr, user.Weight)
	}
	if user.Height != "null" {
		userFieldsArr = append(userFieldsArr, user.Height)
	}
	if user.BMI != "null" {
		userFieldsArr = append(userFieldsArr, user.BMI)
	}
	if user.Body_mass_index_value != "null" {
		userFieldsArr = append(userFieldsArr, user.Body_mass_index_value)
	}
	if user.Blood_group != "null" {
		userFieldsArr = append(userFieldsArr, user.Blood_group)
	}
	if user.Underlying_health_issues != "null" {
		userFieldsArr = append(userFieldsArr, user.Underlying_health_issues)
	}
	if user.Other_health_issues != "null" {
		userFieldsArr = append(userFieldsArr, user.Other_health_issues)
	}
	if user.Images != "null" {
		userFieldsArr = append(userFieldsArr, user.Images)
	}
	if user.User_id != "null" {
		userFieldsArr = append(userFieldsArr, user.User_id)
	}
	return len(userFieldsArr)
}
