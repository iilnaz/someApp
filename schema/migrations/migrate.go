package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Migration struct {
	m *migrate.Migrate
}

func NewMigration(host string,
	port string,
	database string,
	username string,
	password string) (Migration, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password,
		host, port, database)

	m, err := migrate.New("file://schema", dsn)
	if err != nil {
		err = errors.Wrap(err, "can't make a new migrate instance")
	}
	return Migration{m: m}, err
}

func (mg *Migration) Up() error {
	err := mg.m.Up()
	if err != nil {
		err = errors.Wrap(err, "can't do migrate up")
	}
	return err
}

func (mg *Migration) Down() error {
	err := mg.m.Down()
	if err != nil {
		err = errors.Wrap(err, "can't do migrate down")
	}
	return err
}
