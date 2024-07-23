package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyTokens struct {
	PublicKey    *string            `json:"public_key" validate:"required"`
	RefreshToken *[]string          `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       primitive.ObjectID `json:"user_id"`
	Id           primitive.ObjectID `json:"_id"`
}
