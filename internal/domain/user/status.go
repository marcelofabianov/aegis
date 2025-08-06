package user

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/marcelofabianov/fault"
)

type UserLoginStatus string

const (
	UserLoginStatusActive   UserLoginStatus = "active"
	UserLoginStatusInactive UserLoginStatus = "inactive"
	UserLoginStatusBlocked  UserLoginStatus = "blocked"
)

func (s UserLoginStatus) IsActive() bool {
	return s == UserLoginStatusActive
}

func (s UserLoginStatus) IsInactive() bool {
	return s == UserLoginStatusInactive
}

func (s UserLoginStatus) IsBlocked() bool {
	return s == UserLoginStatusBlocked
}

func (s UserLoginStatus) String() string {
	return string(s)
}

func (s UserLoginStatus) IsEmpty() bool {
	return s == ""
}

func (s UserLoginStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *UserLoginStatus) UnmarshalJSON(data []byte) error {
	var status string
	if err := json.Unmarshal(data, &status); err != nil {
		return err
	}
	*s = UserLoginStatus(status)
	return nil
}

func (s UserLoginStatus) Value() (driver.Value, error) {
	if s.IsEmpty() {
		return nil, nil
	}

	return s.String(), nil
}

func (s *UserLoginStatus) Scan(src interface{}) error {
	if src == nil {
		*s = ""
		return nil
	}

	var status string
	switch v := src.(type) {
	case string:
		status = v
	case []byte:
		status = string(v)
	default:
		return fault.New(
			ErrInvalidUserLoginStatus,
			fault.WithCode(fault.Internal),
			fault.WithContext("status", status),
		)
	}

	*s = UserLoginStatus(status)
	return nil
}
