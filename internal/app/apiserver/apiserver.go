package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/astorig/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

func Start(config *Config) (err error) {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := NewServer(store, sessionStore)

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
