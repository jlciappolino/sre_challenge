package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jlciappolino/sre_challenge/apitools/infra"
	"github.com/mercadolibre/sre_challenge/logs/internal"
)

func main() {
	r := gin.Default()

	rdb := infra.NewRedisConn()

	handler := internal.NewPubSubHandler(rdb.Client)

	// r.POST("/pubsub/new", handler.NewTopic)
	// r.POST("/pubsub/subscribe/:topicname", handler.AddConsumer)
	r.POST("/pubsub/send/:topicname", handler.WriteLog)

	r.GET("/check/pubsub/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
		return
	})

	cfg, _ := loadConfig()

	for _, topic := range cfg {
		fmt.Printf("loading topic %s\n", topic.Name)
		handler.AddTopic(topic.Name)
		for _, consumer := range topic.Consumers {
			fmt.Printf("loading consumer %s to %s\n", consumer.URL, topic.Name)
			handler.AddConsumer(topic.Name, consumer)
		}

	}

	r.Run()
}

func loadConfig() (cfg []internal.TopicConfig, err error) {
	path := os.Getenv("config_path")
	configFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(byteValue, &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
