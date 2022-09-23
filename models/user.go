package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone"`
	Is_Baned bool               `json:"is_baned" bson:"is_baned"`
	Ban_Time time.Time          `json:"ban_time" bson:"ban_time"`
	Is_Admin bool               `json:"is_admin" bson:"is_admin"`
}

type UserDB struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password [32]byte           `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Phone    string             `json:"phone" bson:"phone"`
	Is_Baned bool               `json:"is_baned" bson:"is_baned"`
	Ban_Time time.Time          `json:"ban_time" bson:"ban_time"`
	Is_Admin bool               `json:"is_admin" bson:"is_admin"`
}

type AuthenHeaderData struct {
	Authorized bool   `json:"authorized" `
	Username   string `json:"username" `
	PlayerID   string `json:"ulayerID" `
	Is_Addmin  bool   `json:"is_Addmin" `
}
