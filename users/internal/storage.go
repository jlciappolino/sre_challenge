package internal

type Storage interface {
	Get(id string) (*User, error)
	Save(user *User) error
}

type inMemoryStorage struct {
	data map[string]*User
}

func NewInMemoryStorage() *inMemoryStorage {
	return &inMemoryStorage{
		data: map[string]*User{},
	}
}

func (s *inMemoryStorage) Get(id string) (u *User, e error) {
	u = s.data[id]
	return
}

func (s *inMemoryStorage) Save(user *User) (e error) {
	s.data[user.ID] = user
	return
}
