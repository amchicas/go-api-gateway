package domain

import "context"

type Auth struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int64  `json:"role"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
}

type AuthService interface {
	//GetUsers(ctx context.context) ([]*User, error)
	Register(ctx context.Context, auth *Auth) (*Auth, error)
	Login(ctx context.Context, auth *Auth) (string, error)
	//	GetUser(ctx context.context, userId int) (*User, error)
	//	CreateUser(ctx context.context, user *User) (*User, error)
	//	UpdateUser(ctx context.context, userId int, user *User) error
	//	DeleteUser(ctx context.context, userId int) error
}
