package apitools

import "github.com/gin-gonic/gin"

func NewChallengeRouter() *gin.Engine {
	r := gin.Default()

	chaosMiddleware := NewChaoticMiddleware(10)

	r.use(chaosMiddleware.Handle)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, pingResponse)
		return
	})

	return r
}
