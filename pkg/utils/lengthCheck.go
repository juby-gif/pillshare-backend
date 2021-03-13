package utils

import (
	"github.com/juby-gif/pillshare-server/internal/models"
)

func GetLengthOfUserField(user *models.User) int {
	userFieldsArr := []interface{}{}
	if user.FirstName != "" {
		userFieldsArr = append(userFieldsArr, user.FirstName)
	}
	if user.MiddleName != "" {
		userFieldsArr = append(userFieldsArr, user.MiddleName)
	}
	if user.LastName != "" {
		userFieldsArr = append(userFieldsArr, user.LastName)
	}
	if user.Username != "" {
		userFieldsArr = append(userFieldsArr, user.Username)
	}
	if user.Email != "" {
		userFieldsArr = append(userFieldsArr, user.Email)
	}
	if user.Password != "" {
		userFieldsArr = append(userFieldsArr, user.Password)
	}

	userFieldsArr = append(userFieldsArr, user.CheckedStatus)

	if user.Age != "" {
		userFieldsArr = append(userFieldsArr, user.Age)
	}
	if user.Gender != "" {
		userFieldsArr = append(userFieldsArr, user.Gender)
	}
	if user.Dob != "" {
		userFieldsArr = append(userFieldsArr, user.Dob)
	}
	if user.Address != "" {
		userFieldsArr = append(userFieldsArr, user.Address)
	}
	if user.City != "" {
		userFieldsArr = append(userFieldsArr, user.City)
	}
	if user.Country != "" {
		userFieldsArr = append(userFieldsArr, user.Country)
	}
	if user.Province != "" {
		userFieldsArr = append(userFieldsArr, user.Province)
	}
	if user.Zip != "" {
		userFieldsArr = append(userFieldsArr, user.Zip)
	}
	if user.Phone != "" {
		userFieldsArr = append(userFieldsArr, user.Phone)
	}
	if user.Weight != "" {
		userFieldsArr = append(userFieldsArr, user.Weight)
	}
	if user.Height != "" {
		userFieldsArr = append(userFieldsArr, user.Height)
	}
	if user.BMI != "" {
		userFieldsArr = append(userFieldsArr, user.BMI)
	}
	if user.BodyMassIndexValue != "" {
		userFieldsArr = append(userFieldsArr, user.BodyMassIndexValue)
	}
	if user.BloodGroup != "" {
		userFieldsArr = append(userFieldsArr, user.BloodGroup)
	}
	if user.UnderlyingHealthIssues != "" {
		userFieldsArr = append(userFieldsArr, user.UnderlyingHealthIssues)
	}
	if user.OtherHealthIssues != "" {
		userFieldsArr = append(userFieldsArr, user.OtherHealthIssues)
	}
	// if len(user.Images) != 0 {
	// 	userFieldsArr = append(userFieldsArr, user.Images)
	// }
	if user.UserId != "" {
		userFieldsArr = append(userFieldsArr, user.UserId)
	}
	return len(userFieldsArr)
}
