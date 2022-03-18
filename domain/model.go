package domain

import (
	"context"
)

type User struct {
	UUID        string `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Age         string `json:"age"`
	PhoneNumber string `json:"phone_number"`
}

type UserRepo interface {
	InsertData(ctx context.Context, u *User) error
	Get(ctx context.Context, id string) ([]User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
	CloseConnection(ctx context.Context)
	NewConnection(host string, port string, database string, username string, password string) error
}
