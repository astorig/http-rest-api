package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/astorig/http-rest-api/internal/app/store/sqlstore"
)

func Start(config *Config) (err error) {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	srv := NewServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
