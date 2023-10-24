package types

import (
	"time"
)

type User struct {
	ID 			int       `db:"id,omitempty"`
	UserID 		string 	  `db:"user_id,omitempty"`
	Active 		int       `db:"user_active"`
	CreatedAt 	time.Time `db:"created_at"`
	UpdatedAt 	time.Time `db:"updated_at"`

	//stripe stuff soon
	// StripeCustomerID 		string
	// StripeSubscriptionID 	string
	// SubscriptionStatus 		string
	// Plan 					string
}

type SignupParams struct {
	Email 		string `params:"email"`
	FullName 	string `params:"required"`
	Password 	string	`params:"required"`
}

type LoginParams struct {
	Email 		string `params:"email"`
	Password 	string	`params:"required"`
}