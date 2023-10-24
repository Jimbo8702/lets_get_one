package types

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID 			int       `db:"id,omitempty"`
	FirstName 	string    `db:"first_name"`
	LastName 	string    `db:"last_name"`
	Email 		string    `db:"email"`
	Active 		int       `db:"user_active"`
	Password 	string    `db:"password"`
	CreatedAt 	time.Time `db:"created_at"`
	UpdatedAt 	time.Time `db:"updated_at"`
	Token 		Token     `db:"-"`

	//stripe stuff soon
	// StripeCustomerID 		string
	// StripeSubscriptionID 	string
	// SubscriptionStatus 		string
	// Plan 					string
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText)); err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

// resets the users password
func (u *User) ResetPassword(id int, password string) (*User, error) {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}
	u.Password = string(newHash)
	return u, nil
}

type SignupParams struct {
	Email 		string `params:"email"`
	FullName 	string `params:"required"`
	Password 	string	`params:"required"`
}

type LoginParams struct {
	Email 		string 	`params:"email"`
	Password 	string	`params:"required"`
}