package internal

import "encoding/json"

//Item holds item related data
type Item struct {
	ID          string  `json:"id" faker:"-"`
	Description string  `json:"description" faker:"word"`
}

func (s *Item) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Item) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
