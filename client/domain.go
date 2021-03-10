package main

//User holds users related data
type User struct {
	ID     string `json:"id" faker:"-"`
	Type   string `json:"type" faker:"oneof: buyer, seller"`
	Name   string `json:"name" faker:"name"`
	Status string `json:"status" faker:"oneof: active, inactive"`
}

//SoldItem holds item related data
type SoldItem struct {
	ID          string `json:"id" faker:"-"`
	Description string `json:"description" faker:"word"`
}

//Item holds item related data
type Item struct {
	ID       string  `json:"id" faker:"-"`
	Price    float64 `json:"price" faker:"amount"`
	Currency string  `json:"currency" faker:"-"`
}

type Result struct{
	User_id string `json:"user_id"`
	Total float64  `json:"total"`
	Total_usd float64 `json:"sumUSD"`
}