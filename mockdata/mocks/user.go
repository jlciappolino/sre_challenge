package mocks

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/mockdata/domain"
	"strconv"
)
type mockUser struct {
	redis infra.RedisConn
}

func NewMockUsers(redis infra.RedisConn) mockUser {
	return mockUser{redis}
}

func (mock mockUser) Do(externalId int, items []*domain.Item) {
	ctx := context.Background()
	user := generateUserMockData(externalId)
	user.ID = strconv.Itoa(externalId)
	for _,item := range items{
		user.Sold_items = append(user.Sold_items, *item)
	}

	if err := mock.redis.Set(ctx,user.ID, user, 0).Err();
		err != nil {
		fmt.Printf("unable to store user struct into storage due to: %s \n", err)
		return
	}
}

func generateUserMockData(id int) *domain.User {
	user := new(domain.User)
	faker.FakeData(user)
	return user
}