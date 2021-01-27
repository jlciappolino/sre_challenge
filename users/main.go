package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/users/internal"
)

var pingResponse = gin.H{"message": "pong"}

func main() {
	r := apitools.NewChallengeRouter()

	handler := internal.NewUserHandler(internal.NewInMemoryStorage())

	r.GET("/users/:id", handler.Get)

	r.Run()
}
