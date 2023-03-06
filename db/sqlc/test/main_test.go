package sqlc

import (
	"database/sql"
	"log"
	"os"
	"project/e-commerce/db/sqlc"
	"project/e-commerce/initial"
	"testing"

	_ "github.com/lib/pq"
)

var testDB *sql.DB
var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	path := "../../../"
	config, err := initial.LoadingConfig(path)
	if err != nil {
		log.Fatal("Can't load config", err)
	}

	testDB, err = initial.LoadingDatabase(config)	
	if err != nil {
		log.Fatal("Cannot connect to db ,", err)
	}

	testQueries = sqlc.New(testDB)
	os.Exit(m.Run())
}
      