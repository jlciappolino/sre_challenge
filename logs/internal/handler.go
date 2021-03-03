package internal

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

type pubsubHandler struct{
	storage *redis.Client
	TopicNames map[string]map[string]string
}

type Endpoint struct {
	Url   string `json:"url"`
}

type MsgTopic struct{
	Msg        string   `json:"msg"`
	Attributes []string `json:"attributes"`
}

func NewPubSubHandler(s *redis.Client) *pubsubHandler {
	return &pubsubHandler{
		storage: s,
		TopicNames: map[string]map[string]string{},
	}
}

func (ps *pubsubHandler) NewTopic(c *gin.Context){
	ctx := context.Background()

	sUUID := uuid.Must(uuid.NewRandom())

	topicName := sUUID.String()

	if _, exists := ps.TopicNames[topicName]; !exists{
		ps.TopicNames[topicName] = map[string]string{}
	}

	ps.storage.PubSubChannels(ctx,topicName)

	c.JSON(http.StatusCreated, topicName)
	return
}

func (ps *pubsubHandler) AddConsumer (c *gin.Context){
	topicName := c.Param("topicname")
	var endPointPath Endpoint

	if bindErr := c.BindJSON(&endPointPath); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	if !validUrl(endPointPath.Url){
		c.JSON(http.StatusBadRequest, gin.H{"message": "You must to sent a valid endpoint"})
		return
	}

	if _, exists := ps.TopicNames[topicName]; exists{
		if ps.TopicNames[topicName] == nil{
			ps.TopicNames[topicName] = map[string]string{endPointPath.Url:endPointPath.Url}
		}else{
			if _, exists := ps.TopicNames[topicName][endPointPath.Url]; !exists{
				ps.TopicNames[topicName][endPointPath.Url] = endPointPath.Url
			}
		}
		c.JSON(http.StatusCreated, "Added")
		return
	}else{
		c.JSON(http.StatusCreated, "Topic doesnt exists, please create a new topic.")
		return
	}
}

func (ps *pubsubHandler) WriteLog (c *gin.Context){
	var topicMsg MsgTopic
	ctx := context.Background()
	topicName := c.Param("topicname")

	if bindErr := c.BindJSON(&topicMsg); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	//broadcasting the messages to the topic
	if endpoints, exists := ps.TopicNames[topicName];exists {
		if len(endpoints) == 0{
			ps.storage.Publish(ctx,topicName, topicMsg)
		}else{
			for _, endpoint := range endpoints {
				//TODO: post to the consumer
				fmt.Printf("Broadcast to %s -> %s\n", topicName, endpoint)
			}
		}
	}

	c.JSON(http.StatusOK, "Sended")
	return
}

func validUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}



