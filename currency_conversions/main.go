package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"github.com/jlciappolino/sre_challenge/currency_conversions/internal"
)

func main() {
	r := apitools.NewChallengeRouter(1000)

	handler := internal.NewCurrencyController()

	r.GET("/currency_conversions/:id", handler.Get)

	r.GET("/check/currency_conversions/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
		return
	})

	r.Run()
}
