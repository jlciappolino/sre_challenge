package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/users/internal"
)

var pingResponse = gin.H{"message": "pong"}

func main() {
	rand.Seed(time.Now().Unix())

	r := apitools.NewChallengeRouter()

	handler := internal.NewUserHandler(internal.NewInMemoryStorage())

	r.GET("/users/:id", handler.Get)

	r.Run()
}
