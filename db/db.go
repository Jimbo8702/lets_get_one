package db

import (
	"database/sql"
	"fmt"

	upper "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Store struct {
	User UserStore
}

func New(dsn string) (*Store, error) {
	pqsql, err := NewPostgres(dsn)
	if err != nil {
		return nil, err
	}
	us := NewPostgresUserStore(pqsql)
	return &Store{
		User: us,
	}, nil
}

func getInsertID(i upper.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	} 
	return i.(int)
}

func NewPostgres(dsn string) (upper.Session, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	upper, err := postgresql.New(db)
	if err != nil {
		return nil, err
	}
	return upper, nil
}