package user

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/marcelofabianov/fault"
)

type Password string

func (p Password) IsEmpty() bool {
	return p == ""
}

func (p Password) String() string {
	return string(p)
}

func (p Password) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Password) UnmarshalJSON(data []byte) {
	var password string
	if err := json.Unmarshal(data, &password); err != nil {
		return
	}
	*p = Password(password)
	return
}

func (p Password) Value() (driver.Value, error) {
	if p.IsEmpty() {
		return nil, nil
	}

	return p.String(), nil
}

func (p *Password) Scan(src interface{}) error {
	if src == nil {
		*p = ""
		return nil
	}

	var password string
	switch v := src.(type) {
	case string:
		password = v
	case []byte:
		password = string(v)
	default:
		return fault.New(
			ErrInvalidPassword,
			fault.WithCode(fault.Internal),
			fault.WithContext("password", password),
		)
	}

	*p = Password(password)
	return nil
}
