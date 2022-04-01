package main

import (
	"QrCode-Api/configs"
	routers "QrCode-Api/routes"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {
	app = gin.Default()
	configs := configs.GetConfigs()
	listen_port := configs.GetString("WEBSERVER_PORT")

	routers.QRCodeRouter(app)

	app.Run(":" + listen_port)
}
