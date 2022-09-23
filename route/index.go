package route

import (
	version1 "https/route/v1"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func IndexRoute() {

	server := gin.Default()
	server.Use(CORSMiddleware())
	v1 := server.Group("/v1")

	version1.Route_Authen(v1)

	server.Run(":" + os.Getenv("port")) //http
	//	server.RunTLS(":"+os.Getenv("port"), "server.pem", "server.key") //https

}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	}
}
