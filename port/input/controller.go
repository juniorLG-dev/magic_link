package input

import (
	"github.com/gin-gonic/gin"
)

type PortController interface {
	SendEmail(*gin.Context)
	VerifyCode(*gin.Context)
}