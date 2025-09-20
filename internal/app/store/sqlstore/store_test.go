package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "postgres://restapi:restapi@127.0.0.1:63343/restapi_test?sslmode=disable"
	}

	os.Exit(m.Run())
}
