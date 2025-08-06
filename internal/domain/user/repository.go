package user

import "context"

type CreateUserRepository interface {
	ExistsUser(ctx context.Context, input UserExistsInput) (bool, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
}
