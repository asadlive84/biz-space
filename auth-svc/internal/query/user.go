package query

import (
	"database/sql"
	"time"
)

func (q *QueryInit) CreateUser(u User) (User, error) {
	tx, err := q.Db.Beginx()
	if err != nil {
		return User{}, err
	}

	stmt, err := tx.PrepareNamed(`
		INSERT INTO users (
			user_id, user_name, full_name, phone_number, address, password_hash, email, created_at
		) VALUES (
			uuid_generate_v4(), :user_name, :full_name, :phone_number, :address, :password_hash, :email, :created_at
		)
		RETURNING user_id, user_name, full_name, phone_number, address, email, created_at
	`)
	if err != nil {
		tx.Rollback()
		return User{}, err
	}

	u.CreatedAt = time.Now()

	var newUser User
	err = stmt.Get(&newUser, &u)
	if err != nil {
		tx.Rollback()
		return User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (q *QueryInit) GetUserByEmail(email string) (*User, error) {
	var user User
	qc := "SELECT * FROM users WHERE email=$1"
	if err := q.Db.Get(&user, qc, email); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return &user, nil
}
func (q *QueryInit) GetUserByPhone(phone string) (*User, error) {
	var user User
	qc := "SELECT * FROM users WHERE phone=$1"
	if err := q.Db.Get(&user, qc, phone); err != nil {
		if err == sql.ErrNoRows {
			return nil, NotFound
		}
		return nil, err
	}
	return &user, nil
}

func (q *QueryInit) GetUserByID(userID string) (User, error) {

	var user User
	const qc = `SELECT * FROM users WHERE id=$1`
	if err := q.Db.Get(&user, qc, userID); err != nil {
		if err == sql.ErrNoRows {
			return user, NotFound
		}

		return user, err
	}
	return user, nil
}
