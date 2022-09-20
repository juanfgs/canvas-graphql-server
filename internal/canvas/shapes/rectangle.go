package shapes

import (
	"encoding/json"
	"database/sql/driver"
	"errors"
)
type Rectangle struct {
	X int64 `json:"x"`  
	Y int64 `json:"y"`
	Width int64 `json:"width"`
	Height int64 `json:"height"`
	Fill string `json:"fill"`
	Outline string `json:"fill"`
}

func (r Rectangle) Value() (driver.Value, error) {
	return json.Marshal(r)
} 

func (r *Rectangle) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok{
		return errors.New("type assertion failed")
	}
	return json.Unmarshal(b,&r)
}
