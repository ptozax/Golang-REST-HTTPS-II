package v1

import (
	"https/controller"

	"github.com/gin-gonic/gin"
)

func AuthenRoute(route *gin.RouterGroup) {

	route.POST("/login", controller.Login)
	route.POST("/token", controller.GetNewToken)
	route.POST("/signup", controller.Signup)

	route.Use(controller.VerlifyAccess())

	{

		route.POST("/logout", controller.Logout)
	}

}
