package utils

import (
	"fmt"
	"https/repository"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string, playerID string, is_admin bool) (string, string) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["playerid"] = playerID //==================== my data ============
	atClaims["admin"] = is_admin    //==================== my data ============
	atClaims["exp"] = time.Now().Add(time.Minute * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	access_token, _ := at.SignedString([]byte(os.Getenv("TokenKey")))

	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	rtClaims["username"] = username
	rtClaims["playerid"] = playerID //==================== my data ============
	rtClaims["admin"] = is_admin    //==================== my data ============
	rtClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refresh_token, _ := rt.SignedString([]byte(os.Getenv("ReTokenKey")))

	repository.SetToken(username, refresh_token, time.Unix(time.Now().Add(time.Minute*30).Unix(), 0), time.Now())

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
