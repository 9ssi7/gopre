package valobj

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (p UUIDArray) Value() (driver.Value, error) {
	if len(p) == 0 {
		return nil, nil
	}
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}

func (p *UUIDArray) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, p)
	case string:
		return p.Scan([]byte(v))
	default:
		return errors.New("not supported")
	}
}

func (p UUIDArray) ToStringArray() []string {
	var arr []string
	for _, v := range p {
		arr = append(arr, v.String())
	}
	return arr
}
