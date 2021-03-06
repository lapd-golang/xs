package entity

import (
	"time"

	xsvalidator "github.com/zainul/xs/internal/pkg/validator"
)

// User is represented of the user
type User struct {
	Email         string    `json:"email" validate:"required,email"`
	Username      string    `json:"username" validate:"required"`
	Balance       float64   `json:"balance"`
	FirstName     string    `json:"first_name" validate:"required"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number" validate:"required"`
	Point         float64   `json:"point"`
	CitizenID     string    `json:"citizen_id" validate:"required"`
	RefferalCode  string    `json:"refferal_code"`
	AccountNumber string    `json:"account_number"`
	CreatedAt     time.Time `json:"created_at"`
	Status        int64     `json:"status"`
}

// Validate ...
func (u *User) Validate() error {
	return xsvalidator.Validate(u)
}
