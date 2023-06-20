package postgres

import (
	"fmt"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPGConn(dbConf *config.Postgres) (*sqlx.DB, error) {
	dbConnStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.PostgresHost,
		dbConf.PostgresPort,
		dbConf.PostgresUser,
		dbConf.PostgresPass,
		dbConf.PostgresDB,
	)
	db, err := sqlx.Connect("postgres", dbConnStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
