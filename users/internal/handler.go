package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//UsersHandler handle requests
type UsersHandler struct {
	storage Storage
}

func NewUserHandler(s Storage) *UsersHandler {
	return &UsersHandler{
		storage: s,
	}
}

//Get find user by id
func (h *UsersHandler) Get(c *gin.Context) {
	id := c.Param("id")

	u, _ := h.storage.Get(id)

	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, u)

	return
}
