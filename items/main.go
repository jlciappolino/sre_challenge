package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/items/internal"
)

func main() {
	r := apitools.NewChallengeRouter(100)

	rdb := infra.NewRedisConn()
	handler := internal.NewItemHandler(rdb.Client)

	r.GET("/items/:id", handler.Get)
	r.GET("/check/items/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
		return
	})

	r.Run()
}
