package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type SoldItem struct {
	storage *redis.Client
}

func NewSoldItemsHandler(s *redis.Client) *SoldItem {
	return &SoldItem{
		storage: s,
	}
}

//Get find soldItems by userid
func (h *SoldItem) Get(c *gin.Context) {
	ctx := context.Background()
	var user User

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

	if err := user.UnmarshalBinary([]byte(u)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Unable to unmarshal data into the user struct"})
		return
	}

	c.JSON(http.StatusOK, user.Sold_items)

	return
}
