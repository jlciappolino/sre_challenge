package main

import (
	"fmt"
	"math/rand"

	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/mockdata/domain"
	"github.com/mercadolibre/sre_challenge/mockdata/mocks"
)

func main() {
	redis := infra.NewRedisConn()
	mockUsers := mocks.NewMockUsers(redis)
	mockItem := mocks.NewMockItem(redis)

	idItem := 0

	for i := 1; i < 20; i++ {
		items := []*domain.Item{}

		for j := 1; j < rand.Intn(10); j++ {
			idItem++
			item, err := mockItem.Do(idItem)
			if err != nil {
				fmt.Printf("unable to initialize item data due to: %s\n", err.Error())
				continue
			}

			items = append(items, item)
		}

		mockUsers.Do(i, items)

	}
	fmt.Print("========================Finish mock data========================\n")
}
