package user

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/marcelofabianov/fault"

	hash "github.com/marcelofabianov/aegis/internal/platform/hasher"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
	Validate(password string) error
}

type HashedPassword struct {
	value string
}

func NewHashedPassword(value string) (*HashedPassword, error) {
	p := hash.NewPasswordHasher()
	err := p.Validate(value)
	if err != nil {
		return nil, err
	}

	return &HashedPassword{value: value}, nil
}

func (h HashedPassword) String() string {
	return h.value
}

func (h HashedPassword) IsEmpty() bool {
	return h.value == ""
}

func (h HashedPassword) IsValid() bool {
	return !h.IsEmpty()
}

func (h HashedPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h *HashedPassword) UnmarshalJSON(data []byte) error {
	var hashedPassword string
	if err := json.Unmarshal(data, &hashedPassword); err != nil {
		return err
	}

	return nil
}

func (h HashedPassword) Value() (driver.Value, error) {
	if h.IsEmpty() {
		return nil, nil
	}

	return h.String(), nil
}

func (h *HashedPassword) Scan(src interface{}) error {
	if src == nil {
		*h = ""
		return nil
	}

	var hashedPassword string
	switch v := src.(type) {
	case string:
		hashedPassword = v
	case []byte:
		hashedPassword = string(v)
	default:
		return fault.New(
			ErrInvalidHashPassword,
			fault.WithCode(fault.Internal),
			fault.WithContext("hashedPassword", hashedPassword),
		)
	}

	*h = HashedPassword(hashedPassword)
	return nil
}
