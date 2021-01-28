package internal

//User holds users related data
type User struct {
	ID     string `json:"id" faker:"-"`
	Type   string `json:"type" faker:"oneof: buyer, seller"`
	Name   string `json:"name" faker:"name"`
	Status string `json:"status" faker:"oneof: active, inactive"`
}
