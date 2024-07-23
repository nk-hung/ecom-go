package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name     *string   `json:"name" validate:"required, min=2 max=100"`
	Email    *string   `json:"email" validate:"email,required"`
	Password *string   `json:"password" validate:"required, min=6"`
	Status   *string   `json:"status" validate:"required, eq=ACTIVE|eq=INACTIVE"`
	Verify   *string   `json:"verify" validate:"required, default=false"`
	Roles    *[]string `json:"roles" validate:"required, eq=ADMIN|eq=USER|eq=SHOP"`
	// RefreshToken *string            `json:"refresh_token"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	UserId    string             `json:"user_id"`
	ID        primitive.ObjectID `bson:"_id"`
}
