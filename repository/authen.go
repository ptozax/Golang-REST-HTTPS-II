package repository

import (
	"context"
	"crypto/sha256"
	"fmt"
	"https/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func Password_enclip(data models.User) models.UserDB {
	var newUserByte models.UserDB

	newUserByte.ID = data.ID
	newUserByte.Username = data.Username
	newUserByte.Password = sha256.Sum256([]byte(data.Password))
	newUserByte.Email = data.Email
	newUserByte.Phone = data.Phone
	newUserByte.Is_Baned = false
	newUserByte.Ban_Time = time.Date(time.Now().Year()-10, 1, 1, 0, 0, 0, 0, time.UTC)
	newUserByte.Is_Admin = false

	return newUserByte
}

func CheckUserArradySighup(cur *mongo.Cursor, userDB models.UserDB) (bool, error) {

	var data models.UserDB
	var arrData []models.UserDB
	for cur.Next(context.Background()) {
		err := cur.Decode(&data)
		if err != nil {

			return false, err
		}
		arrData = append(arrData, data)
	}

	for _, value := range arrData {
		fmt.Println(value)
		if value.Username == userDB.Username || value.Email == userDB.Email {

			return false, nil
		}
	}

	return true, nil
}
func IsBaned(userDB *models.UserDB) bool {
	if userDB.Is_Baned == true && time.Now().Before(userDB.Ban_Time) {
		return true
	}
	if userDB.Is_Baned == true && time.Now().After(userDB.Ban_Time) {
		userDB.Is_Baned = false
		userDB.Ban_Time = time.Date(time.Now().Year()-10, 1, 1, 0, 0, 0, 0, time.UTC)
		err := PUT_Mongo_One_ByID(userDB.ID, "user", userDB)
		if err != nil {
			return false
		}
	}
	return false
}
