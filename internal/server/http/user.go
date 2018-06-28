package http

import (
	"github.com/fyukiobr/workshop-go/domain"
	"github.com/gin-gonic/gin"
)

func (h *handler) postUser(c *gin.Context) {

}

func (h *handler) getUser(c *gin.Context) {
	userId := c.Param("id")
	user := domain.User{ID: userId, Name: "Fernando"}
	c.JSON(200, user)
}
