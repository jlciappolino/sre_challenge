package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/jlciappolino/sre_challenge/users/internal"
)

func main() {
	r := apitools.NewChallengeRouter(50)

	rdb := infra.NewRedisConn()

	handler := internal.NewUserHandler(rdb.Client)

	r.GET("/users/:id", handler.Get)
	r.GET("/check/users/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
		return
	})

	r.Run()
}
