package internal

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

//UsersHandler handle requests
type UsersHandler struct {
	storage *redis.Client
}

func NewUserHandler(s *redis.Client) *UsersHandler {
	return &UsersHandler{
		storage: s,
	}
}

//Get find user by id
func (h *UsersHandler) Get(c *gin.Context) {
	ctx := context.Background()

	id := c.Param("id")

	u, err := h.storage.Get(ctx, "user-"+id).Result()
	if err != nil {
		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error get data"})
		return
	}

	c.JSON(http.StatusOK, u)

	return
}
