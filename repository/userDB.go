package repository

import (
	"context"
	"crypto/sha256"
	"fmt"
	"https/driver"
	"https/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func User_signupDB(newUser *models.User) bool {

	//========================================= Change pass to [32]Byte ===========================================
	var newUserByte models.UserDB

	newUserByte.ID = newUser.ID
	newUserByte.Username = newUser.Username
	newUserByte.Password = sha256.Sum256([]byte(newUser.Password))
	newUserByte.Email = newUser.Email
	newUserByte.Phone = newUser.Phone
	newUserByte.Is_Baned = false
	newUserByte.Ban_Time = time.Date(time.Now().Year()-10, 1, 1, 0, 0, 0, 0, time.UTC)
	newUserByte.Is_Admin = false

	//============================================================================================================

	client := driver.Connect_db()
	collection := client.Database("test").Collection("user")

	//======================================  check user is have in db  =========================================
	var user models.UserDB
	//err := collection.FindOne(context.TODO(), bson.M{"username": newUser.Username}).Decode(&user)
	err := collection.FindOne(context.TODO(), bson.M{"$or": []bson.M{bson.M{"username": newUserByte.Username}, bson.M{"email": newUserByte.Email}}}).Decode(&user)
	if err == nil {
		return false
	} else {
		if err.Error() != "mongo: no documents in result" {
			return false
		}
	}
	//============================================================================================================

	insertResult, _ := collection.InsertOne(context.TODO(), newUserByte)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	driver.Disconnect_DB(client)
	return true
}

func Put_User(user *models.UserDB) error {

	var userDB models.UserDB
	client := driver.Connect_db()
	collection := client.Database("test").Collection("user")

	err := collection.FindOne(context.TODO(), bson.M{"id": user.ID}).Decode(&userDB)

	if err != nil {

		return err
	}
	_, err = collection.UpdateOne(context.Background(), bson.D{{"_id", user.ID}}, bson.M{"$set": &user})

	if err != nil {

		return err
	}
	driver.Disconnect_DB(client)
	return nil
}

func LoginDB(user *models.User) bool {

	var userDB models.UserDB
	passwordByte := sha256.Sum256([]byte(user.Password)) // string to [32]byte for check

	var logCheck bool

	client := driver.Connect_db()
	collection := client.Database("test").Collection("user")
	err := collection.FindOne(context.TODO(), bson.M{"username": user.Username, "password": passwordByte}).Decode(&userDB)

	user.ID = userDB.ID
	user.Username = userDB.Username
	user.Password = ""
	user.Email = userDB.Email
	user.Phone = userDB.Phone
	user.Is_Baned = userDB.Is_Baned
	user.Ban_Time = userDB.Ban_Time
	user.Is_Admin = userDB.Is_Admin
	fmt.Println(user)

	if userDB.Is_Baned == true && time.Now().Before(userDB.Ban_Time) {
		return false
	}

	//================================reset ban====================================
	if userDB.Is_Baned == true && time.Now().After(userDB.Ban_Time) {
		userDB.Is_Baned = false
		_, err = collection.UpdateOne(context.Background(), bson.D{{"_id", userDB.ID}}, bson.M{"$set": &userDB})
		if err != nil {
			return false
		}
	}
	//================================reset ban====================================
	if err != nil {
		return false

	} else {
		logCheck = true
	}
	driver.Disconnect_DB(client)
	return logCheck
}

func SetToken(id string, token string, ex time.Time, sub time.Time) error {
	client, err := driver.ConnectRedis()

	if err != nil {
		return err
	}
	err = client.Set(id, token, ex.Sub(sub)).Err()
	if err != nil {
		return err
	}

	client.Close()
	return nil
}

func DeleteToken(id string) error {
	client, err := driver.ConnectRedis()
	if err != nil {
		return err
	}

	client.Del(id).Result()
	client.Close()

	return nil
}

func GetToken(id string) (string, error) {
	client, err := driver.ConnectRedis()
	if err != nil {
		return "", err
	}

	val, err := client.Get(id).Result()

	if err != nil {
		return "", err
	}
	client.Close()
	return val, nil
}
