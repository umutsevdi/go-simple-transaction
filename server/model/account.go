package model

import "time"

// An account struct that holds credit cards
type Account struct {
	Id           uint32    `json:"id"`
	Fame         string    `json:"first_name"`
	Lname        string    `json:"last_name"`
	Dob          time.Time `json:"date_of_birth"`
	RegisterDate time.Time `json:"register_date"`

	AccName            string `json:"username"`
	password           string
	lastPassword       string
	PasswordChangeDate time.Time `json:"passwordChangeDate"`
	PrivateKey         string    `json:"private_key"`
	CardId             []uint32  `json:"card_ids"`
}
