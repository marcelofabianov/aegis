package usecase

import (
	"context"

	"github.com/marcelofabianov/aegis/internal/domain/user"
)

type CreateUserUseCase struct {
	createUserRepo user.CreateUserRepository
}

func NewCreateUserUseCase(createUserRepo user.CreateUserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{createUserRepo: createUserRepo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input user.NewUserInput) (*user.NewUserOutput, error) {
	//...

	return nil, nil
}
