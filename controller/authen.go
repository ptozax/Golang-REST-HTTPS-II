package controller

import (
	"https/models"
	"https/repository"
	"https/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		message := models.Invalid_syntax()
		c.JSON(http.StatusBadRequest, message)
		return
	} else {

		if repository.LoginDB(&user) == true {

			//-----------------old----------------------
			at, rt := utils.GenerateToken(user.Username, user.ID.Hex(), user.Is_Admin)
			c.JSON(http.StatusOK, models.Authen{
				Username:     user.Username,
				AccessToken:  at,
				RefreshToken: rt,
			})
			//--------------end old--------------------

		} else {
			message := models.User_not_found()
			c.JSON(http.StatusOK, message)
		}

	}
}

func VerlifyAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {

			c.AbortWithStatus(401)

		} else {
			token, err := utils.Get_tokenJWT(authHeader)
			_, _, _, isaddmin := utils.ParseJson_all_Info(token)

			if !isaddmin {

				if !token.Valid {

					if err.Error() == "Token is expired" {

						//c.AbortWithStatus(401)
						c.AbortWithStatusJSON(403, models.Token_expired())
					} else {
						//c.AbortWithStatus(403)
						c.AbortWithStatusJSON(401, models.Invalid_token())
					}
				}
			}

		}

	}
}

func GetNewToken(c *gin.Context) {

	var rt models.RefreshToken

	err := c.ShouldBindJSON(&rt)

	if err != nil {
		message := models.Invalid_syntax()
		c.JSON(http.StatusOK, message)
	} else {

		token, _ := utils.ValidRefreshToken(rt.RefreshToken)

		if !token.Valid {

			//c.AbortWithStatus(403)
			c.JSON(http.StatusOK, models.Invalid_token())
		} else {

			_, username, playerID, is_addmin := utils.ParseJson_all_Info(token)
			tokenVal, _ := repository.GetToken(username)
			if tokenVal == rt.RefreshToken {
				at, rt := utils.GenerateToken(username, playerID, is_addmin)
				c.JSON(http.StatusOK, models.Token{
					AccessToken:  at,
					RefreshToken: rt,
				})
			} else {
				//c.AbortWithStatus(403)
				c.JSON(http.StatusOK, models.Token_not_found())
			}

		}
	}
}

func Logout(c *gin.Context) {

	var rt models.RefreshToken

	err := c.ShouldBindJSON(&rt)

	if err != nil {
		message := models.Invalid_syntax()
		c.JSON(http.StatusOK, message)
	} else {
		token, err := utils.ValidRefreshToken(rt.RefreshToken)

		if !token.Valid {
			if err.Error() == "Token is expired" {
				c.AbortWithStatus(http.StatusForbidden)
				c.JSON(http.StatusOK, models.Token_expired())
			} else {
				//c.AbortWithStatus(403)
				c.JSON(http.StatusOK, models.Invalid_token())
			}
		} else {

			_, username, _, _ := utils.ParseJson_all_Info(token)
			tokenVal, _ := repository.GetToken(username)
			if tokenVal == rt.RefreshToken {

				repository.DeleteToken(username)
				message := models.Logout_success()
				c.JSON(http.StatusOK, message)

			} else {
				//c.AbortWithStatus(403)
				c.JSON(http.StatusOK, models.Invalid_token())

			}

		}
	}

}
func Signup(c *gin.Context) {

	var newuser models.User

	err := c.ShouldBindJSON(&newuser)

	if err != nil {

		message := models.Invalid_syntax()
		c.JSON(http.StatusOK, message)

	} else {

		isSignupPass := repository.User_signupDB(&newuser)

		if isSignupPass == true {
			message := models.Signup_success()
			c.JSON(http.StatusOK, message)
		} else {
			message := models.Username_Or_Email_AlreadyUsed()
			c.JSON(http.StatusOK, message)
		}

	}

}
