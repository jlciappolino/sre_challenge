package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/mockdata/domain"
	"github.com/mercadolibre/sre_challenge/mockdata/mocks"
)

func main() {
	redis := infra.NewRedisConn()
	mockUsers := mocks.NewMockUsers(redis)
	mockItem := mocks.NewMockItem(redis)

	idItem := 0
	items := []*domain.Item{}

	for j := 1; j < 10; j++ {
		idItem++
		item, err := mockItem.Do(idItem)
		if err != nil {
			fmt.Printf("unable to initialize item data due to: %s\n", err.Error())
			continue
		}

		items = append(items, item)
	}

	for i := 1; i < 20; i++ {
		var itemsToUser []*domain.Item

		rand.Seed(time.Now().UnixNano()+int64(i))
		r := rand.New(rand.NewSource(time.Now().Unix()+int64(i)))
		for _, i := range r.Perm(len(items)) {
			itemsToUser = append(itemsToUser, items[i])
		}

		mockUsers.Do(i, itemsToUser[:rand.Intn(len(itemsToUser))])

	}
	fmt.Print("========================Finish mock data========================\n")
}
