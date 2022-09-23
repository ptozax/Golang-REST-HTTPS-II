package v1

import (
	"https/controller"

	"github.com/gin-gonic/gin"
)

func Route_Authen(route *gin.RouterGroup) {
	route.POST("/signup", controller.POST_Signup)
	route.POST("/login", controller.POST_Login)
	route.POST("/token", controller.POST_RefreshToken)

	route.Use(controller.VerlifyAccess())

	{
		route.POST("/logout", controller.POST_Logout)
	}

}
