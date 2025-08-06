package user

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/marcelofabianov/fault"
)

type Role string

const (
	Admin  Role = "admin"
	Common Role = "common"
	Guest  Role = "guest"
)

func (r Role) IsAdmin() bool {
	return r == Admin
}

func (r Role) IsCommon() bool {
	return r == Common
}

func (r Role) IsGuest() bool {
	return r == Guest
}

func (r Role) IsEmpty() bool {
	return r == ""
}

func (r Role) String() string {
	return string(r)
}

func (r Role) IsValid() bool {
	switch r {
	case Admin, Common, Guest:
		return true
	default:
		return false
	}
}

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Role) UnmarshalJSON(data []byte) error {
	var role string
	if err := json.Unmarshal(data, &role); err != nil {
		return err
	}
	*r = Role(role)
	return nil
}

func (r Role) Value() (driver.Value, error) {
	if r.IsEmpty() {
		return nil, nil
	}

	return r.String(), nil
}

func (r *Role) Scan(src interface{}) error {
	if src == nil {
		*r = ""
		return nil
	}

	var role string
	switch v := src.(type) {
	case string:
		role = v
	case []byte:
		role = string(v)
	default:
		return fault.New(
			ErrInvalidRole,
			fault.WithCode(fault.Internal),
			fault.WithContext("role", role),
		)
	}

	*r = Role(role)
	return nil
}
