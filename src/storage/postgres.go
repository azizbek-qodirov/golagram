package storage

import (
	"api-test/src/config"
	"api-test/src/storage/managers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	PgClient *sql.DB
	UM       *managers.UserManager
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	// #################    POSTGRESQL CONNECTION     ###################### //
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to database pgsql!")

	return &Storage{
		PgClient: db,
		UM:       managers.NewUserManager(db),
	}, nil
}
