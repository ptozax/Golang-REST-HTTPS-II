package utils

import (
	"github.com/dgrijalva/jwt-go"
)

const BEARER_SCHEMA = "Bearer "

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

func ParseJson_all_Info(Token *jwt.Token) (bool, string, string, bool) {
	claims, _ := Token.Claims.(jwt.MapClaims)
	return claims["authorized"].(bool), claims["username"].(string), claims["playerid"].(string), claims["admin"].(bool)
}

func Get_tokenJWT(authHeader string) (*jwt.Token, error) {
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := ValidAccessToken(tokenString)
	if err != nil {
		return token, err
	}
	return token, nil
}

func Get_info_Byauth(authHeader string) (bool, string, string, bool) {
	token, _ := Get_tokenJWT(authHeader)
	authorized, username, playerID, Is_admin := ParseJson_all_Info(token)
	return authorized, username, playerID, Is_admin
}
