package mocks

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/mockdata/domain"
)

type mockItem struct {
	redis infra.RedisConn
}

func NewMockItem(redis infra.RedisConn) mockItem {
	return mockItem{redis}
}

func (mock mockItem) Do(externalId int) (*domain.Item, error) {
	ctx := context.Background()
	item := generateItemMockData(externalId)
	item.ID = strconv.Itoa(externalId)

	if err := mock.redis.Set(ctx, "item-"+item.ID, item, 0).Err(); err != nil {
		return nil, fmt.Errorf("unable to store item struct into storage due to: %s", err)
	}

	fmt.Printf("generated Item: %v", item)

	return item, nil
}

func generateItemMockData(id int) *domain.Item {
	item := new(domain.Item)
	faker.FakeData(item)
	return item
}
