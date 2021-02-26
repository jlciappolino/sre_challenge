package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"net/http"
)

func main() {
	r := apitools.NewChallengeRouter()

	//handler := newHandler TODO

	//r.GET("/items/:id", handler.Get)
	r.GET("/check/items/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK,"pong")
		return
	})

	r.Run()
}
