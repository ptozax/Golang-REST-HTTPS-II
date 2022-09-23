package controller

import (
	"https/models"
	"https/utils"

	"github.com/gin-gonic/gin"
)

func VerlifyAccess() gin.HandlerFunc {
	return func(c *gin.Context) {

		utils.Block{
			Try: func() {

				authHeader := c.GetHeader("Authorization")
				if authHeader == "" {
					utils.Throw("NO AUTH")
				}

				token, err := utils.Token_StringToJWT(authHeader)
				if err != nil {
					if err.Error() != "Token is expired" {
						utils.Throw("TOKEN INVALID")
					}
				}

				auth_DATA := utils.Token_JwtToData(token)
				if !auth_DATA.Is_Addmin {
					if !token.Valid {
						if err.Error() == "Token is expired" {
							utils.Throw("TOKEN EXP")
						} else {
							utils.Throw("TOKEN INVALID")
						}
					}

				}

			},
			Catch: func(e utils.Exception) {
				switch e.(string) {
				case "TOKEN INVALID":
					c.AbortWithStatusJSON(400, models.Token_invalid())
				case "NO AUTH":
					c.AbortWithStatusJSON(401, models.Authentication_nil())
				case "TOKEN EXP":
					c.AbortWithStatusJSON(403, models.Token_expired())
				}
			},
			Finally: func() {

			},
		}.Do()

	}
}
