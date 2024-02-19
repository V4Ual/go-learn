package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Last     string             `json:"last,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
