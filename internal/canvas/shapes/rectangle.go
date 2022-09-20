package shapes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Rectangle struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Fill    string `json:"fill"`
	Outline string `json:"outline"`
}

type RectangleList []*Rectangle

func (r Rectangle) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Rectangle) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Rectangle type assertion failed")
	}
	return json.Unmarshal(b, &r)
}

func (r *RectangleList) Scan(value interface{}) error {

	b, ok := value.([]byte)
	if !ok && value == nil {
		r = &RectangleList{}
		return nil
	}
	if !ok {
		return errors.New("RectangleList type assertion failed")
	}
	return json.Unmarshal(b, &r)
}
