package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type ConsumerConfig struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TopicConfig struct {
	Name      string           `json:"name"`
	Consumers []ConsumerConfig `json:"consumers"`
}

type pubsubHandler struct {
	storage    *redis.Client
	TopicNames map[string]map[string]string
}

type Endpoint struct {
	Url string `json:"url"`
}

type MsgTopic struct {
	Msg      string `json:"msg"`
	Category string `json:"category"`
	Api      string `json:"api"`
}

func NewPubSubHandler(s *redis.Client) *pubsubHandler {
	return &pubsubHandler{
		storage:    s,
		TopicNames: map[string]map[string]string{},
	}
}

func (ps *pubsubHandler) NewTopicHandler(c *gin.Context) {
	sUUID := uuid.Must(uuid.NewRandom())

	topicName := sUUID.String()

	ps.AddTopic(topicName)

	c.JSON(http.StatusCreated, topicName)
	return
}

func (ps *pubsubHandler) AddTopic(topicName string) {
	if _, exists := ps.TopicNames[topicName]; !exists {
		ps.TopicNames[topicName] = map[string]string{}
	}
	ps.storage.PubSubChannels(context.Background(), topicName)
}

func (ps *pubsubHandler) AddConsumerHandler(c *gin.Context) {
	topicName := c.Param("topicname")
	var endPointPath Endpoint

	if bindErr := c.BindJSON(&endPointPath); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	if !validUrl(endPointPath.Url) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You must to sent a valid endpoint"})
		return
	}

	if _, exists := ps.TopicNames[topicName]; exists {
		if ps.TopicNames[topicName] == nil {
			ps.TopicNames[topicName] = map[string]string{endPointPath.Url: endPointPath.Url}
		} else {
			if _, exists := ps.TopicNames[topicName][endPointPath.Url]; !exists {
				ps.TopicNames[topicName][endPointPath.Url] = endPointPath.Url
			}
		}
		c.JSON(http.StatusCreated, "Added")
		return
	} else {
		c.JSON(http.StatusCreated, "Topic doesnt exists, please create a new topic.")
		return
	}
}

func (ps *pubsubHandler) AddConsumer(topicName string, c ConsumerConfig) {
	if _, exists := ps.TopicNames[topicName]; exists {
		if ps.TopicNames[topicName] == nil {
			ps.TopicNames[topicName] = map[string]string{c.URL: c.URL}
		} else {
			if _, exists := ps.TopicNames[topicName][c.URL]; !exists {
				ps.TopicNames[topicName][c.URL] = c.URL
			}
		}
	} else {
		panic("topic not exists")
	}
}

func (ps *pubsubHandler) WriteLog(c *gin.Context) {
	var topicMsg MsgTopic
	ctx := context.Background()
	topicName := c.Param("topicname")

	if bindErr := c.BindJSON(&topicMsg); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}

	//broadcasting the messages to the topic
	if endpoints, exists := ps.TopicNames[topicName]; exists {
		if len(endpoints) == 0 {
			ps.storage.Publish(ctx, topicName, topicMsg)
		} else {
			for _, endpoint := range endpoints {

				fmt.Printf("Broadcast to %s -> %s\n", topicName, endpoint)
				json_data, err := json.Marshal(topicMsg)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "was a problem building the request"})
					return
				}

				resp, err := http.Post(endpoint, "application/json",
					bytes.NewBuffer(json_data))

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "error when post the data to the consumer"})
					return
				}

				var res map[string]interface{}

				json.NewDecoder(resp.Body).Decode(&res)

				if resp.StatusCode == 201 {
					if err != nil {
						c.JSON(resp.StatusCode, "Sended")
						return
					}
				} else {
					if err != nil {
						c.JSON(resp.StatusCode, res)
						return
					}
				}
			}
		}
	}

	c.JSON(http.StatusCreated, "Sended")
	return
}

func validUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
