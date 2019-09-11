package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
*	The model for users
*	Author: Lee Tuan
 */
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Email           string             `json:"email,omitempty"`
	FirstName       string             `json:"firstName,omitempty"`
	LastName        string             `json:"lastName,omitempty"`
	Password        string             `json:"password,omitempty"`
	CreatedAt       int64              `json:"createdAt,omitempty"`
	UpdatedAt       int64              `json:"updatedAt,omitempty"`
}