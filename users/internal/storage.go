package internal

type Storage interface {
	Get(id string) (*User, error)
}

type inMemoryStorage struct {
	data map[string]*User
}

func NewInMemoryStorage() *inMemoryStorage {
	return &inMemoryStorage{}
}

func (s *inMemoryStorage) Get(id string) (u *User, e error) {
	u = s.data[id]
	return
}
