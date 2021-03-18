package apitools

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var pingResponse = gin.H{"message": "pong"}

//NewChallengeRouter returns a router with basic config for the challenge
func NewChallengeRouter(avgRT int) *gin.Engine {
	r := gin.Default()

	rand.Seed(time.Now().Unix())

	chaosMiddleware := NewChaoticMiddleware(8, avgRT)

	r.Use(chaosMiddleware.Handle)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, pingResponse)
		return
	})

	return r
}
