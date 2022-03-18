package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"someApp/domain"
)

type storage struct {
	Conn *pgx.Conn
}

func NewPostgresRepo(Conn *pgx.Conn) domain.UserRepo {
	return &storage{Conn: Conn}
}

func (st *storage) CloseConnection(ctx context.Context) {
	st.Conn.Close(ctx)
}

func (st *storage) NewConnection(host string,
	port string, database string,
	username string, password string) error {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password,
		host, port, database)
	Storage, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to connect to database: %v\n", err))
	}
	st.Conn = Storage
	return err
}

func (st *storage) InsertData(ctx context.Context, u *domain.User) error {
	_, err := st.Conn.Exec(ctx, "INSERT INTO users (id, name, surname, age, phone_number) VALUES ($1, $2, $3, $4, $5)",
		u.UUID, u.Name, u.Surname, u.Age, u.PhoneNumber)

	if err != nil {
		return errors.Wrap(err, "cant insert data")
	}
	return err
}

func (st *storage) Get(ctx context.Context, id string) (*[]domain.User, error) {
	rows, err := st.Conn.Query(ctx, "SELECT id, name, surname, age, phone_number FROM users WHERE id = $1", id)
	if err != nil {
		err = errors.Wrap(err, "cant make select")
	}

	var rowSlice []domain.User

	for rows.Next() {
		var r domain.User
		err = rows.Scan(&r.UUID, &r.Name, &r.Surname, &r.Age, &r.PhoneNumber)
		if err != nil {
			err = errors.Wrap(err, "cant scan")
		}
		rowSlice = append(rowSlice, r)
	}
	return &rowSlice, err
}

func (st *storage) Update(ctx context.Context, u *domain.User) error {
	_, err := st.Conn.Exec(ctx, "UPDATE users SET name = $1, surname = $2, age = $3, phone_number = $4 WHERE id = $5",
		u.Name, u.Surname, u.Age, u.PhoneNumber, u.UUID)

	if err != nil {
		return errors.Wrap(err, "cant update data")
	}
	return err
}

func (st *storage) Delete(ctx context.Context, id string) error {
	_, err := st.Conn.Exec(ctx, "DELETE FROM users WHERE id = $1",
		id)
	if err != nil {
		return errors.Wrap(err, "cant delete data")
	}
	return err
}
