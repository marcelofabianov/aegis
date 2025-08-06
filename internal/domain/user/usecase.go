package user

import "context"

type NewUserOutput struct {
	User *User
}

type CreateUser interface {
	Execute(ctx context.Context, input NewUserInput) (NewUserOutput, error)
}
