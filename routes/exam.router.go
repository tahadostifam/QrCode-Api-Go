package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/tahadostifam/gingo/web_server/controllers"
)

func ExamRouter(router *gin.Engine) {
	router.GET("/", controllers.RootAction)
	router.GET("/makeuser", controllers.MakeUserAction)
}