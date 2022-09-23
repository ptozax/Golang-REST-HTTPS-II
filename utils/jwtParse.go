package utils

import (
	"https/models"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const BEARER_SCHEMA = "Bearer "

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX   New Token   XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

//======================================= Main =====================================================
func Token_JwtToData(token *jwt.Token) models.AuthenHeaderData {
	var authHeaderData models.AuthenHeaderData
	claims, _ := token.Claims.(jwt.MapClaims)

	authHeaderData.Authorized = claims["authorized"].(bool)
	authHeaderData.Username = claims["username"].(string)
	authHeaderData.PlayerID = claims["playerid"].(string)
	authHeaderData.Is_Addmin = claims["admin"].(bool)

	return authHeaderData
}

func Token_TOAH(authString string) string {
	s := string(authString[0:7])
	if strings.Contains(s, BEARER_SCHEMA) {
		authString = authString[len(BEARER_SCHEMA):]
	}
	return authString
}

//==================================================================================================
//======================================= token ====================================================

func Token_StringToJWT(tokenWithAUTH string) (*jwt.Token, error) {
	tokenString := Token_TOAH(tokenWithAUTH)
	token, err := ValidAccessToken(tokenString)
	if err != nil {
		return token, err
	}
	return token, nil
}
func Token_StringToData(tokenWithAUTH string) (models.AuthenHeaderData, error) {
	var tokenData models.AuthenHeaderData
	Tokenjwt, err := Token_StringToJWT(tokenWithAUTH)
	if err != nil {
		return tokenData, err
	}
	tokenData = Token_JwtToData(Tokenjwt)
	return tokenData, nil
}

//==================================================================================================
//======================================= ReFtoken =================================================

func RefToken_StringToJWT(tokenWithAUTH string) (*jwt.Token, error) {
	tokenString := Token_TOAH(tokenWithAUTH)
	token, err := ValidRefreshToken(tokenString)
	if err != nil {
		return token, err
	}
	return token, nil

}

func RefToken_StringToData(tokenWithAUTH string) (models.AuthenHeaderData, error) {
	var tokenData models.AuthenHeaderData
	Tokenjwt, err := RefToken_StringToJWT(tokenWithAUTH)
	if err != nil {
		return tokenData, err
	}
	tokenData = Token_JwtToData(Tokenjwt)
	return tokenData, nil
}

//===========================================================================================
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
//XXXXXXXXXXXXXXXXXXXXXXXXXXX [ Old jwtparse ] XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func ParseJson(Token *jwt.Token) string {
	claims, _ := Token.Claims.(jwt.MapClaims)
	return claims["username"].(string)
}

func ParseJsonPlayerID(Token *jwt.Token) string {
	claims, _ := Token.Claims.(jwt.MapClaims)
	return claims["playerid"].(string)
}

func ParseJson_Is_Addmin(Token *jwt.Token) bool {
	claims, _ := Token.Claims.(jwt.MapClaims)
	return claims["admin"].(bool)
}
