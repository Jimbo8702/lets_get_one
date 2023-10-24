package db

import (
	"time"

	"github.com/Jimbo8702/lets_get_one/types"
	upper "github.com/upper/db/v4"
)

//authenticate token will be moved out of the tokenStore

type TokenStore interface {
	// GetUserForToken(string) (*types.User, error)
	GetTokensForUser(int) ([]*types.Token, error)
	GetById(int) (*types.Token, error)
	GetByToken(string) (*types.Token, error)
	DeleteById(int) error
	DeleteByToken(string) error
	Insert(types.Token, types.User) error
}

// func Cond(key string, data any) upper.Cond {
// 	return upper.Cond{key: data}
// }

type PostgresTokenStore struct {
	collection upper.Collection
}

func NewPostgresTokenStore(up upper.Session) TokenStore {
	return &PostgresTokenStore{
		collection: up.Collection("tokens"), 
	}
}
// Might move this out to handler logic 
// func (pg *PostgresTokenStore) GetUserForToken(token string) (*types.User, error) {
// 	var resultToken types.Token
// 	res := pg.collection.Find(upper.Cond{"token": token})
// 	if err := res.One(&resultToken); err != nil {
// 		return nil, err
// 	}
// 	u, err := pg.us.GetById(resultToken.UserID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	u.Token = resultToken
// 	return u, nil
// }

func (pg *PostgresTokenStore) GetTokensForUser(id int) ([]*types.Token, error) {
	var tokens []*types.Token
	res := pg.collection.Find(upper.Cond{"user_id": id})
	if err := res.All(&tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}

func (pg *PostgresTokenStore) GetById(id int) (*types.Token, error) {
	var token types.Token
	res := pg.collection.Find(upper.Cond{"id": id})
	if err := res.One(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (pg *PostgresTokenStore) GetByToken(plainText string) (*types.Token, error) {
	var token types.Token
	res := pg.collection.Find(upper.Cond{"token": plainText})
	if err := res.One(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

func (pg *PostgresTokenStore) DeleteById(id int) error {
	res := pg.collection.Find(upper.Cond{"id": id})
	if err := res.Delete(); err != nil {
		return err
	}
	return nil
}

func (pg *PostgresTokenStore) DeleteByToken(plainText string) error {
	res := pg.collection.Find(upper.Cond{"token": plainText})
	if err := res.Delete(); err != nil {
		return err
	}
	return nil
}

func (pg *PostgresTokenStore) Insert(token types.Token, u types.User) error {
	res := pg.collection.Find(upper.Cond{"user_id": u.ID})
	if err := res.Delete(); err != nil {
		return err
	}
	// add user data to the token
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	token.FirstName = u.FirstName
	token.Email = u.Email

	_, err := pg.collection.Insert(token)
	if err != nil {
		return err
	}
	return nil
}