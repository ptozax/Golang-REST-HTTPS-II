package utils

import (
	"fmt"
	"https/repository"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string, playerID string, is_admin bool) (string, string) {

	token_exp, _ := strconv.Atoi(os.Getenv("token_EXP"))
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["playerid"] = playerID //==================== my data ============
	atClaims["admin"] = is_admin    //==================== my data ============
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(token_exp)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	access_token, _ := at.SignedString([]byte(os.Getenv("TokenKey")))

	Retoken_exp, _ := strconv.Atoi(os.Getenv("token_EXP"))
	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	rtClaims["username"] = username
	rtClaims["playerid"] = playerID //==================== my data ============
	rtClaims["admin"] = is_admin    //==================== my data ============
	rtClaims["exp"] = time.Now().Add(time.Minute * time.Duration(Retoken_exp)).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refresh_token, _ := rt.SignedString([]byte(os.Getenv("ReTokenKey")))

	repository.POST_Redis_String(username, refresh_token, Retoken_exp)

	return access_token, refresh_token
}

func ValidAccessToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an Error")
			}
			return []byte(os.Getenv("TokenKey")), nil
		})
	return token, err
}

func ValidRefreshToken(accessToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an Error")
			}
			return []byte(os.Getenv("ReTokenKey")), nil
		})
	return token, err
}
