package models

type User struct {
	First_name               string `json:"first_name"`
	Middle_name              string `json:"middle_name"`
	Last_name                string `json:"last_name"`
	Username                 string `json:"username"`
	Email                    string `json:"email"`
	Password                 string `json:"password"`
	Checked_status           bool   `json:"checked_status"`
	Age                      int16  `json:"age"`
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
