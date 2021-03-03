package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/logs/internal"
	"net/http"
)

func main() {
	r := gin.Default()

	rdb :=infra.NewRedisConn()

	handler := internal.NewPubSubHandler(rdb.Client)

	r.POST("/pubsub/new/", handler.NewTopic)
	r.POST("/pubsub/subscribe/:topicname", handler.AddConsumer)
	r.POST("/pubsub/send/:topicname", handler.WriteLog)

	r.GET("/check/pubsub/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK,"pong")
		return
	})

	r.Run()
}
