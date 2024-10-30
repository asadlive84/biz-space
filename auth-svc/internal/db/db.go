package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pressly/goose"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
)

type DBConfig struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string
	POSTGRES_HOST     string
	PORT              string
	JWT_SECRET_KEY    string
}

type DbHandler struct {
	DBMigrate *migrate.Migrate
	SqlDB     *sql.DB
}
func DbInit(dbConfig, configMigrationDir string) (DbHandler, error) {
	var logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	logger.SetOutput(os.Stdout)
	logger.Infof("db url is %+v", dbConfig)
	db, err := sql.Open("postgres", dbConfig)

	if err != nil {
		logger.Errorf("db open issue %+v", err)
		return DbHandler{}, err
	}

	migrationsDir := configMigrationDir
	logger.Infof("Applying migrations from directory: %s", migrationsDir)
	logger.Infof("migrationsDir dir is %+v", migrationsDir)
	if err := goose.Up(db, migrationsDir); err != nil {
		logger.Errorf("goose migration issue %+v", err)
		return DbHandler{}, err
	}

	logger.Infoln("Database migrations applied success!")

	return DbHandler{
		// DBMigrate: m,
		SqlDB: db,
	}, nil
}

func NewDatabaseConfig(config DBConfig) (string, error) {
	var dbParams []string
	dbParams = append(dbParams, fmt.Sprintf("user=%s", config.POSTGRES_USER))
	if host := config.POSTGRES_HOST; host == "" {
		dbParams = append(dbParams, fmt.Sprintf("host=%s", "localhost"))
	} else {
		dbParams = append(dbParams, fmt.Sprintf("host=%s", config.POSTGRES_HOST))
	}
	dbParams = append(dbParams, fmt.Sprintf("port=%s", config.POSTGRES_PORT))
	dbParams = append(dbParams, fmt.Sprintf("dbname=%s", config.POSTGRES_DB))
	if password := config.POSTGRES_PASSWORD; password != "" {
		dbParams = append(dbParams, fmt.Sprintf("password=%s", password))
	}

	dbParams = append(dbParams, fmt.Sprintf("sslmode=%s", "disable"))
	// dbParams = append(dbParams, fmt.Sprintf("sslmode=%s",
	// 	config.GetString("database.sslMode")))
	// dbParams = append(dbParams, fmt.Sprintf("connect_timeout=%d",
	// 	config.GetInt("database.connectionTimeout")))
	// dbParams = append(dbParams, fmt.Sprintf("statement_timeout=%d",
	// 	config.GetInt("database.statementTimeout")))
	// dbParams = append(dbParams, fmt.Sprintf("idle_in_transaction_session_timeout=%d",
	// 	config.GetInt("database.idleInTransactionSessionTimeout")))

	return strings.Join(dbParams, " "), nil
}
