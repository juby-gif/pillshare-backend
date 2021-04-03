// package that handles all the email alerts to corresponding users
package utils

// import (
// 	"github.com/mailgun/mailgun-go/v4"
// )

// func SendSimpleMessage(domain, apiKey string) (string, error) {
	
// 	mg := mailgun.NewMailgun(domain, apiKey, publicApiKey)
// 	m := mg.NewMessage(
// 		"Excited User <mailgun@YOUR_DOMAIN_NAME>",
// 		"Hello",
// 		"Testing some Mailgun awesomeness!",
// 		"YOU@YOUR_DOMAIN_NAME",
// 	)
// 	_, id, err := mg.Send(m)
// 	return id, err
// }