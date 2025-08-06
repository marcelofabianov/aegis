package hash

import (
	"github.com/marcelofabianov/doberman"
	"github.com/marcelofabianov/fault"
)

type PasswordHasher struct{}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (p PasswordHasher) Hash(password string) (string, error) {
	pass := doberman.NewHashedPassword(password)

	return pass.String(), nil
}

func (p PasswordHasher) Compare(hashedPassword, password string) (bool, error) {
	d := doberman.NewArgo2Hasher(nil)

	pass, err := doberman.NewPassword(password)
	if err != nil {
		return false, err
	}

	hashed := doberman.NewHashedPassword(hashedPassword)

	return d.Compare(pass, hashed)
}

func (p PasswordHasher) Validate(password string) error {
	if password == "" {
		return fault.New("password cannot be empty",
			fault.WithCode(fault.Invalid),
			fault.WithContext("field", "password"))
	}

	v := doberman.NewPasswordValidator(nil)

	err := v.Validate(password)

	return err
}
