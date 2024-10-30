package query

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create a new QueryInit instance
	q := QueryInit{
		Db: sqlxDB,
	}

	// Define the input user
	inputUser := User{
		UserName:     "testuser",
		FullName:     "Test User",
		PhoneNumber:  "1234567890",
		Address:      "123 Test St",
		PasswordHash: "hashedpassword",
		Email:        "testuser@example.com",
	}

	// Set up the mock expectations
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO users").
		ExpectQuery().
		WithArgs(
			inputUser.UserName,
			inputUser.FullName,
			inputUser.PhoneNumber,
			inputUser.Address,
			inputUser.PasswordHash,
			inputUser.Email,
			sqlmock.AnyArg(), // created_at
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"user_id", "user_name", "full_name", "phone_number", "address", "email", "created_at",
		}).AddRow(
			"6e11b15c-c75f-4592-a631-757d6a330811", inputUser.UserName, inputUser.FullName, inputUser.PhoneNumber,
			inputUser.Address, inputUser.Email, time.Now(),
		))
	mock.ExpectCommit()

	// Call the CreateUser function
	actualUser, err := q.CreateUser(inputUser)
	assert.NoError(t, err)

	// Assert that the actual user matches the expected user
	assert.Equal(t, inputUser.UserName, actualUser.UserName)
	assert.Equal(t, inputUser.FullName, actualUser.FullName)
	assert.Equal(t, inputUser.PhoneNumber, actualUser.PhoneNumber)
	assert.Equal(t, inputUser.Address, actualUser.Address)
	assert.Equal(t, inputUser.Email, actualUser.Email)
	assert.NotZero(t, actualUser.CreatedAt)

	// Ensure that all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateUser_DatabaseError(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create a new QueryInit instance
	q := QueryInit{
		Db: sqlxDB,
	}

	// Define the input user
	inputUser := User{
		UserName:     "testuser",
		FullName:     "Test User",
		PhoneNumber:  "1234567890",
		Address:      "123 Test St",
		PasswordHash: "hashedpassword",
		Email:        "testuser@example.com",
	}

	// Set up the mock expectations
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO users").
		ExpectQuery().
		WithArgs(
			inputUser.UserName,
			inputUser.FullName,
			inputUser.PhoneNumber,
			inputUser.Address,
			inputUser.PasswordHash,
			inputUser.Email,
			sqlmock.AnyArg(), // created_at
		).
		WillReturnError(sqlmock.ErrCancelled) // Simulate a database error
	mock.ExpectRollback()

	// Call the CreateUser function
	actualUser, err := q.CreateUser(inputUser)
	assert.Error(t, err)
	assert.Equal(t, User{}, actualUser) // Expect an empty User struct on error

	// Ensure that all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreateUser_DuplicateUsername(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Wrap the mock database with sqlx
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Create a new QueryInit instance
	q := QueryInit{
		Db: sqlxDB,
	}

	// Define the input user
	inputUser := User{
		UserName:     "testuser",
		FullName:     "Test User",
		PhoneNumber:  "1234567890",
		Address:      "123 Test St",
		PasswordHash: "hashedpassword",
		Email:        "testuser@example.com",
	}

	// Set up the mock expectations
	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO users").
		ExpectQuery().
		WithArgs(
			inputUser.UserName,
			inputUser.FullName,
			inputUser.PhoneNumber,
			inputUser.Address,
			inputUser.PasswordHash,
			inputUser.Email,
			sqlmock.AnyArg(), // created_at
		).
		WillReturnError(errors.New("pq: duplicate key value violates unique constraint \"users_user_name_key\"")) // Simulate a unique constraint violation
	mock.ExpectRollback()

	// Call the CreateUser function
	actualUser, err := q.CreateUser(inputUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate key value violates unique constraint") // Check if the error contains the unique constraint violation message
	assert.Equal(t, User{}, actualUser)                                               // Expect an empty User struct on error

	// Ensure that all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
