package db

import (
	"time"

	"github.com/Jimbo8702/lets_get_one/types"
	upper "github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	GetAll() ([]*types.User, error)
	GetByEmail(string) (*types.User, error)
	GetById(int) (*types.User, error)
	Update(types.User) error 
	Delete(int) error
	Insert(*types.User) (int, error) 
}

type PostgresUserStore struct {
	collection upper.Collection
}

func NewPostgresUserStore(up upper.Session) UserStore {
	return &PostgresUserStore{
		collection: up.Collection("users"), 
	}
}

func (ps *PostgresUserStore) GetAll() ([]*types.User, error) {
	var all []*types.User
	res := ps.collection.Find().OrderBy("last_name")
	if err := res.All(&all); err != nil {
		return nil, err
	}
	return all, nil
}

func (ps *PostgresUserStore) GetByEmail(email string) (*types.User, error) {
	var u *types.User
	res := ps.collection.Find(upper.Cond{"email":email})
	if err := res.One(&u); err != nil {
		return nil, err
	}
	return u, nil
}

func (ps *PostgresUserStore) GetById(id int) (*types.User, error) {
	var u *types.User
	res := ps.collection.Find(upper.Cond{"id": id})
	if err := res.One(&u); err != nil {
		return nil, err
	}
	return u, nil
} 
 
func (ps *PostgresUserStore) Update(u types.User) error {
	u.UpdatedAt = time.Now()
	res := ps.collection.Find(upper.Cond{"id": u.ID})
	if err := res.Update(&u); err != nil {
		return err
	}
	return nil
}

func (ps *PostgresUserStore) Delete(id int) error {
	res := ps.collection.Find(upper.Cond{"id": id})
	if err := res.Delete(); err != nil {
		return err
	}
	return nil
}

func (ps *PostgresUserStore) Insert(u *types.User) (int, error) {
	newHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return 0, err
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Password = string(newHash)

	res, err := ps.collection.Insert(u)
	if err != nil {
		return 0, err
	}
	id := getInsertID(res.ID())
	return id, nil
} 
