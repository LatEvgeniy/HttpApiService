package main

import (
	"HttpApiService/config"
	"HttpApiService/entrypoints"
	"HttpApiService/processing"
	"HttpApiService/providers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

func main() {
	envCfg := config.GetCfg()
	rabbitProvider := providers.NewRabbitProvider(envCfg)

	orderProcessor := processing.OrderProcessor{RabbitProvider: rabbitProvider}
	userProcessor := processing.UserProcessor{RabbitProvider: rabbitProvider}

	orderEntrypoint := entrypoints.OrderEntrypoint{RabbitProvider: rabbitProvider, OrderPorcessor: &orderProcessor}
	userEntrypoint := entrypoints.UserEntrypoint{RabbitProvider: rabbitProvider, UserProcessor: &userProcessor}

	runServer(&orderEntrypoint, &userEntrypoint)
}

func runServer(orderEntrypoint *entrypoints.OrderEntrypoint, userEntrypoint *entrypoints.UserEntrypoint) {
	router := gin.Default()

	router.GET("/user", userEntrypoint.GetUserById)
	router.POST("/user", userEntrypoint.CreateUser)
	router.PUT("/user", userEntrypoint.AddMoney)

	router.POST("/order", orderEntrypoint.CreateOrder)
	router.GET("/order", orderEntrypoint.GetUserOrders)
	router.DELETE("/order", orderEntrypoint.RemoveOrder)

	router.Run("0.0.0.0:8080")
}
