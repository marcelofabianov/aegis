package user

import (
	"encoding/json"

	"github.com/marcelofabianov/doberman"
	"github.com/marcelofabianov/fault"
	"github.com/marcelofabianov/gobrick/types"
)

type NewUserInput struct {
	Name        string          `json:"name"`
	Phone       string          `json:"phone"`
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	Role        string          `json:"role"`
	Preferences json.RawMessage `json:"preferences,omitempty"`
}

type UserExistsInput struct {
	Email *types.Email
	Phone *types.Phone
}

type FromUserInput struct {
	ID             types.UUID
	Name           string
	Email          types.Email
	Phone          types.Phone
	HashedPassword HashedPassword
	Role           Role
	Status         UserLoginStatus
	Preferences    json.RawMessage
	CreatedAt      types.CreatedAt
	UpdatedAt      types.UpdatedAt
	ArchivedAt     types.ArchivedAt
	DeletedAt      types.DeletedAt
	Version        types.Version
}

type User struct {
	ID             types.UUID       `db:"id" json:"id"`
	Name           string           `db:"name" json:"name"`
	Email          types.Email      `db:"email" json:"email"`
	Phone          types.Phone      `db:"phone" json:"phone"`
	HashedPassword HashedPassword   `db:"hashed_password" json:"-"`
	Role           Role             `db:"role" json:"role"`
	Status         UserLoginStatus  `db:"status" json:"status"`
	Preferences    json.RawMessage  `db:"preferences" json:"preferences"`
	CreatedAt      types.CreatedAt  `db:"created_at" json:"created_at"`
	UpdatedAt      types.UpdatedAt  `db:"updated_at" json:"updated_at"`
	ArchivedAt     types.ArchivedAt `db:"archived_at" json:"archived_at"`
	DeletedAt      types.DeletedAt  `db:"deleted_at" json:"deleted_at"`
	Version        types.Version    `db:"version" json:"version"`
}

func NewUser(input NewUserInput, hasher PasswordHasher) (*User, error) {
	if input.Email == "" {
		return nil, fault.New("email cannot be empty",
			fault.WithCode(fault.Invalid),
			fault.WithContext("field", "email"))
	}
	email, err := types.NewEmail(input.Email)
	if err != nil {
		return nil, fault.Wrap(err, "invalid email format",
			fault.WithCode(fault.Invalid))
	}

	if input.Phone == "" {
		return nil, fault.New("phone cannot be empty",
			fault.WithCode(fault.Invalid),
			fault.WithContext("field", "phone"))
	}
	phone, err := types.NewPhone(input.Phone)
	if err != nil {
		return nil, fault.Wrap(err, "invalid phone format",
			fault.WithCode(fault.Invalid))
	}

	if input.Name == "" {
		return nil, fault.New("name cannot be empty",
			fault.WithCode(fault.Invalid),
			fault.WithContext("field", "name"))
	}

	role := Role(input.Role)
	if !role.IsValid() {
		return nil, fault.New("invalid role",
			fault.WithCode(fault.Invalid),
			fault.WithContext("role", input.Role))
	}

	validator := doberman.NewPasswordValidator(nil)
	if err := validator.Validate(input.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := hasher.Hash(input.Password)
	if err != nil {
		return nil, fault.Wrap(err, "failed to hash password",
			fault.WithCode(fault.Internal))
	}

	id, err := types.NewUUID()
	if err != nil {
		return nil, fault.Wrap(err, "failed to generate UUID", fault.WithCode(fault.Internal))
	}

	return &User{
		ID:             id,
		Name:           input.Name,
		Email:          email,
		Phone:          phone,
		HashedPassword: hashedPassword,
		Role:           role,
		Status:         UserLoginStatusActive,
		Preferences:    input.Preferences,
		CreatedAt:      types.NewCreatedAt(),
		UpdatedAt:      types.NewUpdatedAt(),
		Version:        types.Version(1),
	}, nil
}

func FromUser(input FromUserInput) *User {
	return &User{
		ID:             input.ID,
		Name:           input.Name,
		Email:          input.Email,
		Phone:          input.Phone,
		HashedPassword: input.HashedPassword,
		Role:           input.Role,
		Status:         input.Status,
		Preferences:    input.Preferences,
		CreatedAt:      input.CreatedAt,
		UpdatedAt:      input.UpdatedAt,
		ArchivedAt:     input.ArchivedAt,
		DeletedAt:      input.DeletedAt,
		Version:        input.Version,
	}
}
