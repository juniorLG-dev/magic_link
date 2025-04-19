package routes

import (
	"magic_link/port/input"

	"github.com/gin-gonic/gin"
)

func InitRoutes(rg *gin.RouterGroup, ct input.PortController) {
	rg.POST("/user", ct.SendEmail)
	rg.GET("/verify/:code", ct.VerifyCode)
}