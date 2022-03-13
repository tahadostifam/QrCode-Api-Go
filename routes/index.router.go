package routers

import (
	"github.com/gin-gonic/gin"
	controllers "$1/controllers"
)

func IndexRouter(router *gin.Engine) {
	router.GET("/", controllers.RootAction)
	router.GET("/makeuser", controllers.MakeUserAction) // Testing Model
}