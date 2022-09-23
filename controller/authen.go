package controller

import (
	"https/models"
	"https/repository"
	"https/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX SIGN UP XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func POST_Signup(c *gin.Context) {
	utils.Block{
		Try: func() {

			var user models.User
			err := c.ShouldBindJSON(&user)
			if err != nil {
				utils.Throw("INVALID DATA")
			}

			UserDB := repository.Password_enclip(user)
			UserDB.Is_Baned = false
			UserDB.Ban_Time = time.Date(time.Now().Year()-10, 1, 1, 0, 0, 0, 0, time.UTC)
			UserDB.Is_Admin = false
			filter := bson.M{"$or": []bson.M{bson.M{"username": UserDB.Username}, bson.M{"email": UserDB.Email}}}
			cur, err := repository.GET_Mongo_Many("user", filter)
			if err != nil {
				utils.Throw("GET DATA ERROR")
			}

			signupPass, err := repository.CheckUserArradySighup(cur, UserDB)
			if err != nil {
				utils.Throw("DECODE DATA ERROR")
			}
			if !signupPass {
				utils.Throw("USERNAME OR EMAIL ALREADY USED")
			}
			err = repository.POST_Mongo_One("user", UserDB)
			if err != nil {
				utils.Throw("INSERT DATA ERROR")

			}

			c.JSON(http.StatusOK, models.Signup_success())

		},
		Catch: func(e utils.Exception) {
			switch e.(string) {
			case "INVALID DATA":
				c.AbortWithStatusJSON(400, models.Invalid_syntax())
			case "GET DATA ERROR":
				c.AbortWithStatusJSON(500, models.Get_Data_Error())
			case "DECODE DATA ERROR":
				c.AbortWithStatusJSON(500, models.Decode_error())
			case "USERNAME OR EMAIL ALREADY USED":
				c.AbortWithStatusJSON(403, models.Username_Or_Email_AlreadyUsed())
			case "INSERT DATA ERROR":
				c.AbortWithStatusJSON(500, models.Insert_error())
			default:
				c.AbortWithStatusJSON(400, models.Signup_error())
			}
		},
		Finally: func() {

		},
	}.Do()
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX LOGIN XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func POST_Login(c *gin.Context) {
	utils.Block{
		Try: func() {

			var user models.User
			err := c.ShouldBindJSON(&user)
			if err != nil {
				utils.Throw("INVALID DATA")
			}
			UserDB := repository.Password_enclip(user)
			var DbFound models.UserDB
			//filter := bson.M{"username": UserDB.Username, "password": UserDB.Password}
			filter := bson.M{"$and": []bson.M{bson.M{"username": UserDB.Username}, bson.M{"password": UserDB.Password}}}
			err = repository.GET_Mongo_FindOne_withFilter("user", filter, &DbFound)
			if err != nil {
				if err.Error() == "mongo: no documents in result" {
					utils.Throw("WRONG USERNAME OR PASSWORD")
				} else if err != nil {
					utils.Throw("GET DATA ERROR")
				}
			}

			if repository.IsBaned(&DbFound) {
				utils.Throw("HAS BEEN BANED")
			}

			at, rt := utils.GenerateToken(DbFound.Username, DbFound.ID.Hex(), DbFound.Is_Admin)
			c.JSON(http.StatusOK, models.Authen{
				Username:     DbFound.Username,
				AccessToken:  at,
				RefreshToken: rt,
			})
		},
		Catch: func(e utils.Exception) {
			switch e.(string) {
			case "INVALID DATA":
				c.AbortWithStatusJSON(400, models.Invalid_syntax())
			case "GET DATA ERROR":
				c.AbortWithStatusJSON(500, models.Get_Data_Error())
			case "DECODE DATA ERROR":
				c.AbortWithStatusJSON(500, models.Decode_error())
			case "WRONG USERNAME OR PASSWORD":
				c.AbortWithStatusJSON(403, models.Incorrect_Username_or_Password())
			case "HAS BEEN BANED":
				c.AbortWithStatusJSON(403, models.Baned_User())
			default:
				c.AbortWithStatusJSON(400, models.Login_error())
			}
		},
		Finally: func() {

		},
	}.Do()

}

func POST_Logout(c *gin.Context) {
	utils.Block{
		Try: func() {

			var retoken models.RefreshToken
			err := c.ShouldBindJSON(&retoken)

			if err != nil {
				utils.Throw("INVALID DATA")
			}
			//token
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				utils.Throw("NO AUTH")
			}
			tokenData, err := utils.Token_StringToData(authHeader)
			if err != nil {
				utils.Throw("TOKEN INVALID")
			}
			//ref Token
			refTokenData, err := utils.RefToken_StringToData(retoken.RefreshToken)
			if err != nil {
				utils.Throw("INVALID DATA")
			}
			if tokenData.Username != refTokenData.Username || tokenData.PlayerID != refTokenData.PlayerID {
				utils.Throw("TOKEN NOT MATCH")
			}
			stringToken, err := repository.GET_Redis_String(refTokenData.Username)
			if err != nil {
				utils.Throw("TOKEN NOT FOUND")
			}
			if stringToken != retoken.RefreshToken {
				utils.Throw("TOKEN NOT MATCH")
			}
			err = repository.DEL_Redis_String(tokenData.Username)
			if err != nil {
				utils.Throw("LOGOUT ERROR")
			}

			c.JSON(http.StatusOK, models.Logout_success())
		},
		Catch: func(e utils.Exception) {
			switch e.(string) {
			case "INVALID DATA":
				c.AbortWithStatusJSON(400, models.Invalid_syntax())
			case "TOKEN INVALID":
				c.AbortWithStatusJSON(400, models.Token_invalid())
			case "NO AUTH":
				c.AbortWithStatusJSON(401, models.Authentication_nil())
			case "TOKEN NOT MATCH":
				c.AbortWithStatusJSON(401, models.Token_not_match())
			case "TOKEN NOT FOUND":
				c.AbortWithStatusJSON(400, models.Token_not_found())

			default:
				c.AbortWithStatusJSON(400, models.Logout_error())
			}
		},
		Finally: func() {

		},
	}.Do()

}
func POST_RefreshToken(c *gin.Context) {

	utils.Block{
		Try: func() {

			var retoken models.RefreshToken
			err := c.ShouldBindJSON(&retoken)

			if err != nil {
				utils.Throw("INVALID DATA")
			}
			//ref Token
			refTokenData, err := utils.RefToken_StringToData(retoken.RefreshToken)
			if err != nil {
				utils.Throw("TOKEN INVALID")
			}
			//IMDB reset token
			stringRETOKEN, err := repository.GET_Redis_String(refTokenData.Username)
			if err != nil {
				utils.Throw("TOKEN NOT FOUND")
			}
			if stringRETOKEN != retoken.RefreshToken {
				utils.Throw("TOKEN NOT MATCH")
			}
			rETOKENDATA, err := utils.RefToken_StringToData(stringRETOKEN)
			if err != nil {
				utils.Throw("TOKEN INVALID")
			}
			at, rt := utils.GenerateToken(rETOKENDATA.Username, rETOKENDATA.PlayerID, rETOKENDATA.Is_Addmin)
			c.JSON(http.StatusOK, models.Token{AccessToken: at, RefreshToken: rt})
		},
		Catch: func(e utils.Exception) {
			switch e.(string) {
			case "INVALID DATA":
				c.AbortWithStatusJSON(400, models.Invalid_syntax())
			case "TOKEN INVALID":
				c.AbortWithStatusJSON(400, models.Token_invalid())
			case "TOKEN NOT FOUND":
				c.AbortWithStatusJSON(400, models.Token_not_found())
			case "TOKEN NOT MATCH":
				c.AbortWithStatusJSON(401, models.Token_not_match())
			default:
				c.AbortWithStatusJSON(400, models.Get_Data_Error())
			}

		},
		Finally: func() {

		},
	}.Do()

}
