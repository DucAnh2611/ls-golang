package main

import (
	"strings"

	"github.com/DucAnh2611/ls-golang/config"
	"github.com/DucAnh2611/ls-golang/logging"
	"github.com/DucAnh2611/ls-golang/middlewares"
	"github.com/DucAnh2611/ls-golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	logging.InitLogger("app.log")
	config.LoadEnv()

	router := gin.Default()

	// Register middlewares
	router.Use(middlewares.LoggerMiddleware())

	// Register routes
	routes.RegisterRoutes(router)

	router.Run(strings.Join([]string{":", config.GetEnv("PORT", "2611")}, ""))
}
