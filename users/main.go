package main

import (
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/users/internal"
)

var pingResponse = gin.H{"message": "pong"}

func main() {
	r := apitools.NewChallengeRouter()

	storage := internal.NewInMemoryStorage()

	handler := internal.NewUserHandler(storage)

	buildData(storage)

	r.GET("/users/:id", handler.Get)

	r.Run()
}

func buildData(storage internal.Storage) {
	for i := 1; i < 301; i++ {
		user := new(internal.User)
		faker.FakeData(user)
		user.ID = strconv.Itoa(i)
		storage.Save(user)
	}
}
