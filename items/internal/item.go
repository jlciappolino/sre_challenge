package internal

import (
	"encoding/json"
)

//Item holds item related data
type Item struct {
	ID          string  `json:"id" faker:"-"`
	Price       float64 `json:"price" faker:"amount"`
	Currency	string	`json:"currency" faker:"-"`
}

func (s *Item) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Item) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
