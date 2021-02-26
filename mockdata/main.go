package main

import (
	"fmt"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/mockdata/domain"
	"github.com/mercadolibre/sre_challenge/mockdata/mocks"
	"math/rand"
)

func main() {
	redis := infra.NewRedisConn()
	mockUsers := mocks.NewMockUsers(redis)
	mockItem := mocks.NewMockItem(redis)

	idItem := 0

	for i := 1; i < 2; i++ {
		var items []*domain.Item

		for j := 1; j < rand.Intn(10); j++ {
			idItem++
			item, err := mockItem.Do(idItem)
			if err != nil{
				fmt.Errorf("unable to initialize item data due to:", err)
				continue
			}

			items = append(items, item)
		}

		mockUsers.Do(i, items)

	}

}
