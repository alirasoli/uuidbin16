package uuidbin16

import (
	"database/sql/driver"
	"github.com/google/uuid"
)

type UUIDBin16 uuid.UUID

func StringToBin16(s string) (UUIDBin16, error) {
	if s == "" {
		return UUIDBin16(uuid.Nil), nil
	}
	u, err := uuid.Parse(s)
	if err != nil {
		return UUIDBin16{}, err
	}
	return UUIDBin16(u), nil
}

func (u UUIDBin16) String() string {
	return uuid.UUID(u).String()
}

func (u UUIDBin16) GormType() string {
	return "binary(16)"
}

func (u UUIDBin16) MarshalJSON() ([]byte, error) {
	str := "\"" + u.String() + "\""
	return []byte(str), nil
}

func (u *UUIDBin16) UnmarshalJSON(b []byte) error {
	s, err := uuid.ParseBytes(b)
	*u = UUIDBin16(s)
	return err
}

func (u *UUIDBin16) Scan(src interface{}) error {
	if src == nil {
		*u = UUIDBin16(uuid.Nil)
		return nil
	}
	bytes := src.([]byte)
	parseBytes, err := uuid.FromBytes(bytes)
	*u = UUIDBin16(parseBytes)
	return err
}

func (u UUIDBin16) Value() (driver.Value, error) {
	if u == UUIDBin16(uuid.Nil) {
		return nil, nil
	}
	return uuid.UUID(u).MarshalBinary()
}
