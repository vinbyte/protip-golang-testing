package domain

import "context"

type UserUsecase interface {
	GetAll(context.Context) (code int, response interface{})
	GetByID(context.Context, int) (code int, response interface{})
}

type UserRepository interface {
	Fetch(ctx context.Context, userID int) ([]User, error)
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserList struct {
	Items []User `json:"items"`
}
