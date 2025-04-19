package main

import (
	"magic_link/adapter/output/smtp"
	"magic_link/adapter/output/cache"
	"magic_link/adapter/input/controller"
	"magic_link/adapter/input/routes"
	"magic_link/application/service"

	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"

	"log"
	"os"
)

func main() {
	router := gin.Default()
	sender := os.Getenv("SENDER")
	pass := os.Getenv("PASS")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	smtp := smtp.NewSenderCredentials(sender, pass)
	cache := cache.NewCache(client)
	service := service.NewService(cache, smtp)
	controller := controller.NewController(service)
	routes.InitRoutes(&router.RouterGroup, controller)

	log.Panic(router.Run(":8080"))
}