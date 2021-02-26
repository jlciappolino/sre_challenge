package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools"
	"net/http"
)

func main() {
	check := gin.New()

	r := apitools.NewChallengeRouter()

	//handler := newHandler TODO

	//r.GET("/items/:id", handler.Get)
	check.GET("/items/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK,"pong")
		return
	})

	r.Run()
	check.Run(":8081")

}
