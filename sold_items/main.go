package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/sold_items/internal"
)

func main() {
	r := apitools.NewChallengeRouter(50)

	rdb := infra.NewRedisConn()

	handler := internal.NewSoldItemsHandler(rdb.Client)

	r.GET("/sold_items/:id", handler.Get)
	r.GET("/check/sold_items/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
		return
	})

	r.Run()

}
