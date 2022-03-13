package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tahadostifam/gingo/web_server/configs"
	routers "github.com/tahadostifam/gingo/web_server/routes"
)

var app *gin.Engine

func main() {
	app = gin.Default()
	configs := configs.GetConfigs()
	listen_port := configs.GetString("WEBSERVER_PORT")

	routers.ExamRouter(app)

	app.Run(":" + listen_port)
}