package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
<<<<<<< HEAD
	"someApp/internal/user"
=======
>>>>>>> a2c8e2f (lets try)
)

type Storage struct {
	conn *pgx.Conn
}

<<<<<<< HEAD
func (st *Storage) CloseConnection() {
=======
func (st Storage) CloseConnection() {
>>>>>>> a2c8e2f (lets try)
	st.conn.Close(context.Background())
}

func NewConnection(host string,
	port string,
	database string,
	username string,
	password string) (Storage, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password,
		host, port, database)
<<<<<<< HEAD

	storage, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("unable to connect to database: %v\n", err))
=======
	storage, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		errors.Wrap(err, fmt.Sprintf("unable to connect to database: %v\n", err))
>>>>>>> a2c8e2f (lets try)
	}
	return Storage{conn: storage}, err
}

<<<<<<< HEAD
func (st *Storage) InsertData(u user.User) error {
	_, err := st.conn.Exec(context.Background(), "INSERT INTO users (id, name, surname, age, phone_number) VALUES ($1, $2, $3, $4, $5)",
		u.UUID, u.Name, u.Surname, u.Age, u.PhoneNumber)
	if err != nil {
		return errors.Wrap(err, "cant insert data")
	}
	return err
}

func (st *Storage) Get(id string) ([]user.User, error) {

	rows, err := st.conn.Query(context.Background(), "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		err = errors.Wrap(err, "cant make select")
	}

	var rowSlice []user.User

	for rows.Next() {
		var r user.User
		err = rows.Scan(&r.UUID, &r.Name, &r.Surname, &r.Age, &r.PhoneNumber)
		if err != nil {
			err = errors.Wrap(err, "cant scan")
		}
		rowSlice = append(rowSlice, r)
	}
	return rowSlice, err
}

func (st *Storage) Update(u user.User) error {
	_, err := st.conn.Exec(context.Background(), "UPDATE users SET name = $1, surname = $2, age = $3, phone_number = $4 WHERE id = $5",
		u.Name, u.Surname, u.Age, u.PhoneNumber, u.UUID)
	if err != nil {
		return errors.Wrap(err, "cant update data")
	}
	return err
}

func (st *Storage) Delete(id string) error {
	_, err := st.conn.Exec(context.Background(), "DELETE FROM users WHERE id = $1",
		id)
	if err != nil {
		return errors.Wrap(err, "cant delete data")
	}
=======
func InsertData(storage Storage, name string, age int) error {
	_, err := storage.conn.Exec(context.Background(), "INSERT INTO accounts (username, age) VALUES ($1, $2)", name, age)
	if err != nil {
		return errors.Wrap(err, "cannot create connection")
	}
	fmt.Println("Data inserted")
>>>>>>> a2c8e2f (lets try)
	return err
}
