package controller

import (
	"magic_link/port/input"
	"magic_link/adapter/input/model/request"

	"github.com/gin-gonic/gin"

	"net/http"
)

type controller struct {
	service input.PortService
}

func NewController(service input.PortService) *controller {
	return &controller{
		service: service,
	}
}

func (ct *controller) SendEmail(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := ct.service.SendEmail(userRequest.Email); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.Status(http.StatusOK)
}

func (ct *controller) VerifyCode(c *gin.Context) {
	code := c.Param("code")

	if err := ct.service.VerifyCode(code); err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email verified successfully"})
}