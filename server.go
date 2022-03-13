package main

import (
	"$1/configs"
	routers "$1/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {
	app = gin.Default()
	configs := configs.GetConfigs()
	listen_port := configs.GetString("WEBSERVER_PORT")

	routers.ExamRouter(app)

	app.Run(":" + listen_port)
}
