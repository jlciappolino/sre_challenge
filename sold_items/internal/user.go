package internal

import "encoding/json"

//User holds users related data
type User struct {
	ID     string `json:"id" faker:"-"`
	Type   string `json:"type" faker:"oneof: buyer, seller"`
	Name   string `json:"name" faker:"name"`
	Status string `json:"status" faker:"oneof: active, inactive"`
	Sold_items []Item `json:"sold_items" faker:"-"`
}

func (s *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}