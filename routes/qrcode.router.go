package routers

import (
	controllers "QrCode-Api/controllers"

	"github.com/gin-gonic/gin"
)

func QRCodeRouter(router *gin.Engine) {
	router.GET("/new", controllers.MakeQRCodeAction)
}
