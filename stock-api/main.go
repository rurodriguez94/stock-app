package main

import (
	"net/http"
	"os"

	"github.com/stock-app/stock-api/internal/pkg/logs"

	"github.com/stock-app/stock-api/internal/web/middleware"

	"github.com/stock-app/stock-api/internal/stock/handler"
	"github.com/stock-app/stock-api/internal/stock/provider"
	"github.com/stock-app/stock-api/internal/stock/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	logs.InitDefault()

	alphaAPIKey := os.Getenv("ALPHA_APIKEY")
	if alphaAPIKey == "" {
		alphaAPIKey = "demo"
	}

	router.Use(middleware.ErrorHandler)

	sh := handler.NewStockHandler(service.NewStockService(
		provider.NewStockProvider(alphaAPIKey),
		provider.NewSecurityProvider(http.Client{}),
	))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/stock/:symbol", sh.FetchData)
	router.POST("/stock/decrypt/:token", sh.DecryptStockData)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
