package internal

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type ItemHandler struct {
	storage *redis.Client
}

func NewItemHandler(s *redis.Client) *ItemHandler {
	return &ItemHandler{
		storage: s,
	}
}

//Get find Items by userid
func (h *ItemHandler) Get(c *gin.Context) {
	ctx := context.Background()
	var item Item
	id := c.Param("id")

	u, err := h.storage.Get(ctx, "item-"+id).Result()
	if err != nil {
		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error get data"})
		return
	}

	if err := item.UnmarshalBinary([]byte(u)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Unable to unmarshal data into the item struct"})
		return
	}
	
	c.JSON(http.StatusOK, item)

	return
}
