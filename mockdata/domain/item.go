package domain

//Item holds item related data
type Item struct {
	ID         string `json:"id" faker:"-"`
	Description string `json:"description" faker:"description"`
	Price float64  `json:"price" faker:"price"`
}
