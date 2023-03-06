package initial

import (
	"database/sql"
	"errors"
	"fmt"
	"project/e-commerce/config"

	_ "github.com/lib/pq"
)

type databaseMaker struct {
	Config config.Config
}

func LoadingDatabase(config config.Config) (*sql.DB, error) {
	make := &databaseMaker{
		Config: config,
	}
	sqlDb, err := sql.Open(config.Db.DBDriver, make.psqlDsn())
	if err != nil {
		err := errors.New("Can't open database: " + err.Error())
		return nil, err
	}
	sqlDb.SetMaxIdleConns(config.Db.MaxIdle)
	sqlDb.SetMaxOpenConns(config.Db.MaxOpen)
	return sqlDb, nil
}


func (d *databaseMaker) psqlDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		d.Config.Db.Host,
		d.Config.Db.Port,
		d.Config.Db.Username,
		d.Config.Db.Db,
		d.Config.Db.Password,
	)
}