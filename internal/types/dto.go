package types

import "encoding/json"

type Dto interface {
	Marshal()
	Unmarshal()
}

type Data[T any] struct {
	dto T
}

func (c *Data[T]) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(c.dto)
	if err != nil {
		return bytes, err
	}

	return bytes, nil
}
