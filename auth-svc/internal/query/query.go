package query

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type QueryInit struct {
	Db *sqlx.DB
}

var NotFound = errors.New("not found")

type User struct {
	UserID       string       `db:"user_id"`
	UserName     string       `db:"user_name"`
	FullName     string       `db:"full_name"`
	PhoneNumber  string       `db:"phone_number"`
	Address      string       `db:"address"`
	PasswordHash string       `db:"password_hash"`
	Email        string       `db:"email"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}

type Query interface {
	CreateUser(u User) (User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByPhone(phone string) (*User, error)
	GetUserByID(userID string) (User, error)
}
