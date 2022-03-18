package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func NewConnection(host string,
	port string,
	database string,
	username string,
	password string) (*pgx.Conn, error) {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password,
		host, port, database)
	storage, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unable to connect to database: %v\n", err))
	}
	return storage, nil
}
